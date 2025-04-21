package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Product represents a catalog item
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Cache represents a simple in-memory cache with expiration
type Cache struct {
	sync.RWMutex
	items    map[string]cacheItem
	hits     atomic.Int64
	misses   atomic.Int64
	cleanupInterval time.Duration
}

type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// Metrics tracks server performance metrics
type Metrics struct {
	requestCount   atomic.Int64
	responseTimeNs atomic.Int64
	requestsInFlight atomic.Int64
}

// Server encapsulates the HTTP server and its dependencies
type Server struct {
	products []Product
	cache    *Cache
	metrics  *Metrics
	bufPool  *sync.Pool
}

// NewCache creates a new cache instance with automatic cleanup
func NewCache(cleanupInterval time.Duration) *Cache {
	c := &Cache{
		items:           make(map[string]cacheItem),
		cleanupInterval: cleanupInterval,
	}
	go c.startCleanup()
	return c
}

func (c *Cache) startCleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	for range ticker.C {
		c.cleanup()
	}
}

func (c *Cache) cleanup() {
	c.Lock()
	defer c.Unlock()
	now := time.Now()
	for k, v := range c.items {
		if now.After(v.expiration) {
			delete(c.items, k)
		}
	}
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.items[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(expiration),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, exists := c.items[key]
	if !exists {
		c.misses.Add(1)
		return nil, false
	}
	if time.Now().After(item.expiration) {
		c.misses.Add(1)
		return nil, false
	}
	c.hits.Add(1)
	return item.value, true
}

// gzipWriter wraps http.ResponseWriter to provide gzip compression
type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// compressionMiddleware adds gzip compression to responses
func compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz := gzip.NewWriter(w)
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

// metricsMiddleware collects request metrics
func (s *Server) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.metrics.requestCount.Add(1)
		s.metrics.requestsInFlight.Add(1)
		defer s.metrics.requestsInFlight.Add(-1)

		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		s.metrics.responseTimeNs.Add(duration.Nanoseconds())
	})
}

func (s *Server) handleProducts(w http.ResponseWriter, r *http.Request) {
	// Try to get from cache first
	if cached, ok := s.cache.Get("products"); ok {
		if products, ok := cached.([]byte); ok {
			w.Header().Set("Content-Type", "application/json")
			w.Write(products)
			return
		}
	}

	// Get buffer from pool
	buf := s.bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer s.bufPool.Put(buf)

	// Encode products to JSON using the buffer
	if err := json.NewEncoder(buf).Encode(s.products); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Cache the encoded result
	s.cache.Set("products", buf.Bytes(), 5*time.Minute)

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := map[string]interface{}{
		"total_requests":      s.metrics.requestCount.Load(),
		"requests_in_flight": s.metrics.requestsInFlight.Load(),
		"avg_response_time_ms": float64(s.metrics.responseTimeNs.Load()) / float64(s.metrics.requestCount.Load()) / 1e6,
		"cache_hits":          s.cache.hits.Load(),
		"cache_misses":        s.cache.misses.Load(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	// Initialize server components
	server := &Server{
		products: []Product{
			{ID: 1, Name: "Product 1", Description: "Description 1", Price: 19.99},
			{ID: 2, Name: "Product 2", Description: "Description 2", Price: 29.99},
			{ID: 3, Name: "Product 3", Description: "Description 3", Price: 39.99},
		},
		cache: NewCache(1 * time.Minute),
		metrics: &Metrics{},
		bufPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	// Create router and add routes
	mux := http.NewServeMux()
	mux.Handle("/products", server.metricsMiddleware(compressionMiddleware(http.HandlerFunc(server.handleProducts))))
	mux.Handle("/metrics", server.metricsMiddleware(http.HandlerFunc(server.handleMetrics)))

	// Configure the HTTP server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
} 