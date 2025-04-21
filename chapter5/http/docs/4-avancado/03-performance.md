# Performance e Otimização em Go HTTP

## Introdução

A performance é um aspecto crítico em aplicações web. Go oferece várias ferramentas e técnicas para otimizar o desempenho de servidores HTTP. Vamos explorar as principais estratégias de otimização.

## 1. Configuração do Servidor

### 1.1 Timeouts Apropriados

```go
server := &http.Server{
    ReadTimeout:       5 * time.Second,    // Tempo para ler request
    WriteTimeout:      10 * time.Second,   // Tempo para escrever response
    IdleTimeout:       120 * time.Second,  // Tempo máximo de conexão ociosa
    ReadHeaderTimeout: 2 * time.Second,    // Tempo para ler headers
    MaxHeaderBytes:    1 << 20,            // 1MB máximo para headers
}
```

### 1.2 Keep-Alive

```go
server := &http.Server{
    IdleTimeout: 120 * time.Second,  // Tempo máximo de keep-alive
    Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Connection", "keep-alive")
        w.Header().Set("Keep-Alive", "timeout=120")
        // ... handler code ...
    }),
}
```

## 2. Otimização de Handlers

### 2.1 Pooling de Objetos

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Obter buffer do pool
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    // Usar buffer para processamento
    json.NewEncoder(buf).Encode(data)
    w.Write(buf.Bytes())
}
```

### 2.2 Streaming de Dados

```go
func streamHandler(w http.ResponseWriter, r *http.Request) {
    // Habilitar streaming
    w.Header().Set("Transfer-Encoding", "chunked")
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming not supported", http.StatusInternalServerError)
        return
    }

    // Enviar dados em chunks
    for data := range dataChannel {
        fmt.Fprintf(w, "%s\n", data)
        flusher.Flush()
    }
}
```

## 3. Compressão

### 3.1 Middleware de Compressão

```go
func compressionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verificar se cliente aceita gzip
        if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
            next.ServeHTTP(w, r)
            return
        }

        // Criar writer com gzip
        gz := gzip.NewWriter(w)
        defer gz.Close()

        w.Header().Set("Content-Encoding", "gzip")
        next.ServeHTTP(gzipResponseWriter{
            ResponseWriter: w,
            Writer:        gz,
        }, r)
    })
}

type gzipResponseWriter struct {
    http.ResponseWriter
    io.Writer
}

func (g gzipResponseWriter) Write(data []byte) (int, error) {
    return g.Writer.Write(data)
}
```

## 4. Caching

### 4.1 Cache em Memória

```go
type Cache struct {
    sync.RWMutex
    items map[string]cacheItem
}

type cacheItem struct {
    value      interface{}
    expiration time.Time
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.RLock()
    defer c.RUnlock()

    item, exists := c.items[key]
    if !exists {
        return nil, false
    }

    if time.Now().After(item.expiration) {
        return nil, false
    }

    return item.value, true
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.Lock()
    defer c.Unlock()

    c.items[key] = cacheItem{
        value:      value,
        expiration: time.Now().Add(duration),
    }
}
```

### 4.2 HTTP Caching

```go
func cacheableHandler(w http.ResponseWriter, r *http.Request) {
    // Verificar ETag
    etag := calculateETag(data)
    w.Header().Set("ETag", etag)
    
    if match := r.Header.Get("If-None-Match"); match == etag {
        w.WriteHeader(http.StatusNotModified)
        return
    }

    // Configurar cache
    w.Header().Set("Cache-Control", "public, max-age=300")
    w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
    
    // Enviar dados
    json.NewEncoder(w).Encode(data)
}
```

## 5. Otimização de Banco de Dados

### 5.1 Connection Pool

```go
db, err := sql.Open("postgres", dsn)
if err != nil {
    log.Fatal(err)
}

// Configurar pool
db.SetMaxOpenConns(25)      // Máximo de conexões abertas
db.SetMaxIdleConns(25)      // Máximo de conexões ociosas
db.SetConnMaxLifetime(5 * time.Minute)  // Tempo máximo de vida da conexão
```

### 5.2 Batch Operations

```go
func batchInsert(users []User) error {
    // Preparar query
    stmt, err := db.Prepare(`
        INSERT INTO users (name, email)
        VALUES ($1, $2)
    `)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Executar em batch
    for _, user := range users {
        if _, err := stmt.Exec(user.Name, user.Email); err != nil {
            return err
        }
    }

    return nil
}
```

## 6. Monitoramento de Performance

### 6.1 Métricas Básicas

```go
type Metrics struct {
    requestCount   int64
    responseTime   time.Duration
    errorCount     int64
    activeRequests int64
}

func metricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        atomic.AddInt64(&metrics.activeRequests, 1)
        defer atomic.AddInt64(&metrics.activeRequests, -1)

        start := time.Now()
        next.ServeHTTP(w, r)
        
        atomic.AddInt64(&metrics.requestCount, 1)
        atomic.AddInt64(&metrics.responseTime, time.Since(start).Nanoseconds())
    })
}
```

### 6.2 Profiling

```go
import _ "net/http/pprof"

func main() {
    // Endpoint para profiling
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // ... resto do código ...
}

// Usar: go tool pprof http://localhost:6060/debug/pprof/heap
```

## 7. Boas Práticas

1. **Otimização Prematura**
   - Profile primeiro, otimize depois
   - Use benchmarks para medir melhorias
   - Mantenha o código legível

2. **Concorrência**
   - Use goroutines com cuidado
   - Evite contenção de locks
   - Considere channels para coordenação

3. **Memória**
   - Reutilize objetos quando possível
   - Evite alocações desnecessárias
   - Use slices pré-alocados

4. **I/O**
   - Use buffers apropriados
   - Implemente streaming quando possível
   - Considere compressão para respostas grandes

## Próximos Passos

- [Testes de Carga](04-testes-carga.md)
- [Monitoramento](05-monitoramento.md)
- [Escalabilidade](06-escalabilidade.md) 