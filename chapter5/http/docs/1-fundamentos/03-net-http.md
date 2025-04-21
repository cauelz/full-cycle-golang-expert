# O Pacote net/http em Go

## Visão Geral

O pacote `net/http` é a implementação nativa de HTTP em Go, oferecendo tanto cliente quanto servidor HTTP. É um dos pacotes mais bem projetados da biblioteca padrão, combinando simplicidade com poder e flexibilidade.

## 1. Principais Componentes

### 1.1 Servidor
```go
type Server struct {
    Addr           string        // endereço TCP para escutar
    Handler        Handler       // handler para servir requisições
    ReadTimeout    time.Duration // timeout para ler requisição inteira
    WriteTimeout   time.Duration // timeout para escrever resposta
    MaxHeaderBytes int          // tamanho máximo do header
    TLSConfig      *tls.Config  // configuração opcional de TLS
    // ... outros campos
}
```

### 1.2 Handler Interface
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

### 1.3 ResponseWriter Interface
```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

### 1.4 Request Struct
```go
type Request struct {
    Method string
    URL *url.URL
    Header Header
    Body io.ReadCloser
    ContentLength int64
    Host string
    RemoteAddr string
    // ... outros campos
}
```

## 2. Criando um Servidor HTTP

### 2.1 Forma Mais Simples
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
    http.ListenAndServe(":8080", nil)
}
```

### 2.2 Com Configuração Customizada
```go
package main

import (
    "net/http"
    "time"
)

func main() {
    server := &http.Server{
        Addr:         ":8080",
        Handler:      nil, // usa DefaultServeMux
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }
    server.ListenAndServe()
}
```

## 3. Handlers e ServeMux

### 3.1 Handler Customizado
```go
type userHandler struct {
    users map[string]User
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(h.users)
    case "POST":
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        // ... salvar usuário
    }
}
```

### 3.2 ServeMux (Roteador)
```go
func main() {
    mux := http.NewServeMux()
    
    // Handler função
    mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Olá!")
    })
    
    // Handler struct
    userHandler := &userHandler{
        users: make(map[string]User),
    }
    mux.Handle("/api/users", userHandler)
    
    http.ListenAndServe(":8080", mux)
}
```

## 4. Cliente HTTP

### 4.1 Cliente Simples
```go
resp, err := http.Get("http://exemplo.com")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
```

### 4.2 Cliente Customizado
```go
client := &http.Client{
    Timeout: time.Second * 10,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
        IdleConnTimeout:     90 * time.Second,
    },
}

req, err := http.NewRequest("GET", "http://exemplo.com", nil)
req.Header.Add("User-Agent", "MeuApp/1.0")

resp, err := client.Do(req)
```

## 5. Middlewares

### 5.1 Padrão de Middleware
```go
func logging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

### 5.2 Usando Middleware
```go
func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Olá!")
    })
    
    http.Handle("/", logging(handler))
    http.ListenAndServe(":8080", nil)
}
```

## 6. Recursos Avançados

### 6.1 Servindo Arquivos Estáticos
```go
// Serve diretório /static na raiz
fs := http.FileServer(http.Dir("static"))
http.Handle("/", fs)

// Serve diretório /static em /files/
http.Handle("/files/", http.StripPrefix("/files/", fs))
```

### 6.2 Usando HTTPS
```go
// Com certificados
http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)

// Com Let's Encrypt automático
import "golang.org/x/crypto/acme/autocert"

m := &autocert.Manager{
    Cache:      autocert.DirCache("cert-cache"),
    Prompt:     autocert.AcceptTOS,
    HostPolicy: autocert.HostWhitelist("exemplo.com"),
}

s := &http.Server{
    Addr:      ":https",
    Handler:   mux,
    TLSConfig: m.TLSConfig(),
}

s.ListenAndServeTLS("", "") // certificados vêm do Let's Encrypt
```

### 6.3 Graceful Shutdown
```go
srv := &http.Server{Addr: ":8080"}

go func() {
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

// Espera sinal para shutdown
quit := make(chan os.Signal, 1)
signal.Notify(quit, os.Interrupt)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
    log.Fatal(err)
}
```

## 7. Boas Práticas

1. **Sempre feche o Body das respostas**
   ```go
   resp, err := http.Get(url)
   if err != nil {
       return err
   }
   defer resp.Body.Close()
   ```

2. **Use timeouts apropriados**
   ```go
   srv := &http.Server{
       ReadTimeout:  5 * time.Second,
       WriteTimeout: 10 * time.Second,
       IdleTimeout:  120 * time.Second,
   }
   ```

3. **Limite o tamanho dos payloads**
   ```go
   const maxBodySize = 1 << 20 // 1 MB
   
   func handler(w http.ResponseWriter, r *http.Request) {
       r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
   }
   ```

4. **Use context para cancelamento**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   
   req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
   ```

## Próximos Passos

- [Criando Servidores HTTP](../2-servidor/01-criando-servidor.md)
- [Handlers e ServeMux](../2-servidor/02-handlers-mux.md)
- [Ciclo de Vida de uma Requisição](../2-servidor/03-ciclo-de-vida.md) 