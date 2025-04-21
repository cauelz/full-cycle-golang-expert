# Middlewares em Go

## Introdução

Middlewares são funções que processam requisições HTTP antes ou depois do handler principal. Eles são essenciais para:
- Logging
- Autenticação
- CORS
- Rate Limiting
- Compressão
- Métricas
- Recuperação de pânico

## 1. Estrutura Básica

### 1.1 Assinatura de Middleware

```go
type Middleware func(http.Handler) http.Handler
```

### 1.2 Exemplo Simples

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Chamar próximo handler
        next.ServeHTTP(w, r)
        
        // Log após a execução
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    })
}
```

## 2. Tipos de Middleware

### 2.1 Pré-processamento

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verificar token
        token := r.Header.Get("Authorization")
        if !isValidToken(token) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Token válido, continuar
        next.ServeHTTP(w, r)
    })
}
```

### 2.2 Pós-processamento

```go
func headerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Chamar handler primeiro
        next.ServeHTTP(w, r)
        
        // Adicionar headers após
        w.Header().Set("X-Version", "1.0.0")
    })
}
```

### 2.3 Wrapper Completo

```go
func metricMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Antes
        start := time.Now()
        
        // Wrapper do ResponseWriter para capturar o status
        wrapper := newResponseWriter(w)
        
        // Handler
        next.ServeHTTP(wrapper, r)
        
        // Depois
        duration := time.Since(start)
        log.Printf(
            "status=%d duration=%s method=%s path=%s",
            wrapper.status,
            duration,
            r.Method,
            r.URL.Path,
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    status int
    written bool
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
    return &responseWriter{ResponseWriter: w}
}

func (w *responseWriter) WriteHeader(status int) {
    w.status = status
    w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(b []byte) (int, error) {
    if !w.written {
        w.status = http.StatusOK
    }
    return w.ResponseWriter.Write(b)
}
```

## 3. Encadeamento de Middlewares

### 3.1 Função de Encadeamento

```go
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}
```

### 3.2 Uso

```go
func main() {
    mux := http.NewServeMux()
    
    // Handlers
    mux.HandleFunc("/api/users", handleUsers)
    
    // Middlewares em ordem
    handler := Chain(mux,
        recoveryMiddleware,
        loggingMiddleware,
        authMiddleware,
        corsMiddleware,
    )
    
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

## 4. Middlewares Comuns

### 4.1 CORS

```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Configurar CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", 
            "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Authorization")
            
        // Responder OPTIONS diretamente
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

### 4.2 Rate Limiting

```go
func rateLimitMiddleware(next http.Handler) http.Handler {
    limiter := rate.NewLimiter(rate.Every(time.Second), 100)
    
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", 
                http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### 4.3 Recuperação de Pânico

```go
func recoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Log do stack trace
                stack := debug.Stack()
                log.Printf("panic: %v\n%s", err, stack)
                
                // Resposta de erro
                http.Error(w, "Internal Server Error", 
                    http.StatusInternalServerError)
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

## 5. Boas Práticas

1. **Ordem dos Middlewares**
   ```go
   handler := Chain(mux,
       // 1. Recuperação (mais externo)
       recoveryMiddleware,
       // 2. Logging
       loggingMiddleware,
       // 3. Segurança
       securityHeadersMiddleware,
       // 4. CORS
       corsMiddleware,
       // 5. Autenticação
       authMiddleware,
       // 6. Rate Limiting (mais interno)
       rateLimitMiddleware,
   )
   ```

2. **Contexto para Dados**
   ```go
   func authMiddleware(next http.Handler) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           user, err := authenticate(r)
           if err != nil {
               http.Error(w, "Unauthorized", http.StatusUnauthorized)
               return
           }
           
           // Adicionar ao contexto
           ctx := context.WithValue(r.Context(), "user", user)
           next.ServeHTTP(w, r.WithContext(ctx))
       })
   }
   ```

3. **Configuração via Closure**
   ```go
   func timeoutMiddleware(timeout time.Duration) Middleware {
       return func(next http.Handler) http.Handler {
           return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
               ctx, cancel := context.WithTimeout(r.Context(), timeout)
               defer cancel()
               
               next.ServeHTTP(w, r.WithContext(ctx))
           })
       }
   }
   ```

## 6. Exemplo Completo

```go
type Server struct {
    router     *http.ServeMux
    middleware []Middleware
}

func NewServer(middleware ...Middleware) *Server {
    return &Server{
        router:     http.NewServeMux(),
        middleware: middleware,
    }
}

func (s *Server) Handle(pattern string, handler http.Handler) {
    // Aplicar middlewares específicos da rota
    h := Chain(handler, s.middleware...)
    s.router.Handle(pattern, h)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.router.ServeHTTP(w, r)
}

func main() {
    server := NewServer(
        recoveryMiddleware,
        loggingMiddleware,
        corsMiddleware,
    )
    
    server.Handle("/api/users", handleUsers())
    server.Handle("/api/products", handleProducts())
    
    log.Fatal(http.ListenAndServe(":8080", server))
}
```

## Próximos Passos

- [Segurança e TLS](02-seguranca.md)
- [Performance e Otimização](03-performance.md)
- [Testes de Middleware](04-testes.md) 