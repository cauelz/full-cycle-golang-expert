package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

// Task representa uma tarefa
type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// TaskStore é um armazenamento em memória para tarefas
type TaskStore struct {
	sync.RWMutex
	tasks map[string]Task
}

// NewTaskStore cria um novo armazenamento de tarefas
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[string]Task),
	}
}

// Add adiciona uma nova tarefa
func (s *TaskStore) Add(task Task) {
	s.Lock()
	defer s.Unlock()
	s.tasks[task.ID] = task
}

// Get retorna uma tarefa pelo ID
func (s *TaskStore) Get(id string) (Task, bool) {
	s.RLock()
	defer s.RUnlock()
	task, ok := s.tasks[id]
	return task, ok
}

// List retorna todas as tarefas
func (s *TaskStore) List() []Task {
	s.RLock()
	defer s.RUnlock()
	tasks := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// Middleware é uma função que processa requisições HTTP
type Middleware func(http.Handler) http.Handler

// Chain encadeia múltiplos middlewares
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// LoggingMiddleware registra informações sobre a requisição
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Criar wrapper para capturar o status
		ww := &responseWriter{ResponseWriter: w}

		// Processar requisição
		next.ServeHTTP(ww, r)

		// Log após processamento
		log.Printf(
			"method=%s path=%s status=%d duration=%s",
			r.Method,
			r.URL.Path,
			ww.status,
			time.Since(start),
		)
	})
}

// RecoveryMiddleware recupera de pânicos
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log do stack trace
				log.Printf("panic: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TimeoutMiddleware adiciona timeout para requisições
func TimeoutMiddleware(timeout time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			done := make(chan bool)
			go func() {
				next.ServeHTTP(w, r.WithContext(ctx))
				done <- true
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				w.WriteHeader(http.StatusGatewayTimeout)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "timeout",
				})
			}
		})
	}
}

// CORSMiddleware adiciona headers CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

// TaskHandler gerencia as requisições relacionadas a tarefas
type TaskHandler struct {
	store *TaskStore
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		tasks := h.store.List()
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.ID = fmt.Sprintf("task_%d", time.Now().UnixNano())
		task.CreatedAt = time.Now()
		h.store.Add(task)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Criar store e handler
	store := NewTaskStore()
	handler := &TaskHandler{store: store}

	// Criar mux e adicionar rota
	mux := http.NewServeMux()
	mux.Handle("/tasks", handler)

	// Aplicar middlewares
	finalHandler := Chain(mux,
		RecoveryMiddleware,
		LoggingMiddleware,
		TimeoutMiddleware(5*time.Second),
		CORSMiddleware,
	)

	// Configurar servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: finalHandler,
	}

	// Canal para sinais de término
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor em goroutine
	go func() {
		log.Printf("Servidor iniciado em http://localhost%s", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguardar sinal de término
	<-stop
	log.Println("Desligando servidor...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Erro ao desligar servidor: %v", err)
	}

	log.Println("Servidor desligado")
}