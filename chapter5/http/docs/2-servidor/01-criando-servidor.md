# Criando Servidores HTTP em Go

## Introdução

Em Go, criar um servidor HTTP é surpreendentemente simples graças ao pacote `net/http`. Vamos explorar diferentes formas de criar e configurar servidores HTTP, desde o mais básico até configurações mais avançadas.

## 1. Servidor HTTP Básico

### 1.1 O Servidor Mais Simples
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Olá, Mundo!")
    })
    
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

Este exemplo demonstra:
- Uso do `DefaultServeMux` (quando `nil` é passado para `ListenAndServe`)
- Handler função simples
- Sem configurações especiais

### 1.2 Com Tratamento de Erros
```go
package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    
    log.Println("Iniciando servidor em :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Erro ao iniciar servidor:", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s", r.Method, r.URL.Path)
    fmt.Fprintf(w, "Olá, Mundo!")
}
```

## 2. Servidor com Configurações Customizadas

### 2.1 Usando http.Server
```go
package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    server := &http.Server{
        Addr:         ":8080",
        Handler:      http.HandlerFunc(handler),
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    log.Fatal(server.ListenAndServe())
}
```

### 2.2 Com TLS (HTTPS)
```go
package main

import (
    "log"
    "net/http"
)

func main() {
    server := &http.Server{
        Addr:    ":443",
        Handler: http.HandlerFunc(handler),
    }
    
    log.Fatal(server.ListenAndServeTLS("server.crt", "server.key"))
}
```

## 3. Servidor com Múltiplas Rotas

```go
package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    // Criar novo ServeMux
    mux := http.NewServeMux()
    
    // Rota para API
    mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "Olá, Mundo!",
        })
    })
    
    // Rota para arquivos estáticos
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Iniciar servidor
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    
    server.ListenAndServe()
}
```

## 4. Servidor com Graceful Shutdown

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // Criar servidor
    server := &http.Server{
        Addr:    ":8080",
        Handler: http.HandlerFunc(handler),
    }
    
    // Canal para sinais do sistema
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
    
    // Iniciar servidor em goroutine
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("Erro ao iniciar servidor:", err)
        }
    }()
    log.Print("Servidor iniciado")
    
    // Aguardar sinal de término
    <-done
    log.Print("Servidor está encerrando...")
    
    // Criar contexto com timeout para shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Tentar shutdown gracioso
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Servidor forçado a encerrar:", err)
    }
    
    log.Print("Servidor encerrado com sucesso")
}
```

## 5. Servidor com Middlewares

```go
package main

import (
    "log"
    "net/http"
    "time"
)

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    })
}

// Middleware de CORS
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

func main() {
    // Handler principal
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Olá, Mundo!"))
    })
    
    // Aplicar middlewares
    handler = loggingMiddleware(handler)
    handler = corsMiddleware(handler)
    
    // Criar e iniciar servidor
    server := &http.Server{
        Addr:    ":8080",
        Handler: handler,
    }
    
    log.Fatal(server.ListenAndServe())
}
```

## 6. Boas Práticas

1. **Sempre use timeouts**
   ```go
   server := &http.Server{
       ReadTimeout:       15 * time.Second,
       WriteTimeout:      15 * time.Second,
       IdleTimeout:       60 * time.Second,
       ReadHeaderTimeout: 5 * time.Second,
   }
   ```

2. **Implemente graceful shutdown**
   - Permite que requisições em andamento terminem
   - Fecha conexões de forma limpa
   - Evita perda de dados

3. **Use logging apropriado**
   - Registre erros e eventos importantes
   - Inclua informações úteis para debug
   - Considere usar um formato estruturado (JSON)

4. **Configure CORS corretamente**
   - Seja específico com origens permitidas
   - Limite métodos e headers permitidos
   - Não use `*` em produção

5. **Limite tamanho de payloads**
   ```go
   server.MaxHeaderBytes = 1 << 20 // 1MB
   ```

## Próximos Passos

- [Handlers e ServeMux](02-handlers-mux.md)
- [Ciclo de Vida de uma Requisição](03-ciclo-de-vida.md)
- [Middlewares](../4-avancado/01-middlewares.md) 