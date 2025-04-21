### Ciclo de vida de uma requisição HTTP no **net/http** do Go — passo a passo

1. **Criação (ou reuso) da conexão TCP**  
   - O cliente resolve o DNS e abre uma conexão TCP com `Server.Addr` (normalmente porta :80 ou :443).  
   - Se o cabeçalho `Connection: keep-alive` estava presente em requisições anteriores, a mesma conexão pode ser reaproveitada a partir do *idle connection pool* interno do cliente/servidor.

2. **`net.Listener.Accept()` → goroutine da conexão**  
   - O método `srv.Serve(l net.Listener)` aceita a conexão e imediatamente cria **uma goroutine dedicada** só para aquele *socket*.  
   - Antes de processar qualquer coisa, o servidor dispara (se configurado) o *hook* `ConnState`, informando que o socket está em estado `StateNew`.

3. **Leitura e *parsing* do request**  
   - A goroutine lê o *stream* TCP, decodifica a camada HTTP (HTTP/1.1 ou HTTP/2) e popula um objeto `*http.Request`.  
   - Enquanto faz o *parsing*, ela atualiza `ConnState` para `StateActive`.

4. **Criação do `ResponseWriter` e *contexto***  
   - Para cada requisição, o servidor prepara um `http.response` interno que implementa a interface `http.ResponseWriter`.  
   - O campo `Request.Context()` recebe um `context.Context` com *deadline*, cancelamento e valores (útil em middlewares).

5. **Despacho para o *handler***  
   - O servidor chama `handler.ServeHTTP(w, r)` **em uma nova goroutine** (uma por request).  
   - O *handler* (pode ser o `http.DefaultServeMux` ou qualquer cadeia de middlewares) executa sua lógica de negócio, escreve cabeçalhos e corpo via `w`.

6. **Envio da resposta**  
   - A primeira chamada a `w.WriteHeader` (ou `w.Write`) faz o servidor montar e enviar os cabeçalhos.  
   - Escritas subsequentes vão direto para o *socket* (podendo ser *chunked* ou com `Content-Length`).  
   - Quando o *handler* retorna, o servidor garante que tudo foi *flushed* e chama `ConnState` → `StateIdle`.

7. **Keep‑Alive, *timeouts* e fechamento**  
   - Se o cliente declarou `Connection: keep-alive`, a conexão fica no pool até `Server.IdleTimeout` ou `Server.ReadTimeout` expirar.  
   - Caso contrário, ou se houve erro, o estado vira `StateClosed` e o *socket* é encerrado.

```
Cliente ──TCP──> http.Server ──goroutine──> Handler
   ▲                                          │
   └────────────── HTTP response ‹────────────┘
```

> **Resumo mental**: *socket* aceito → goroutine da conexão → *parsing* → nova goroutine do handler → handler retorna → resposta enviada/flush → decide manter ou fechar.

---

### Exemplo didático completo

```go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Olá, %s!\n", r.URL.Path[1:])
		}),
		ConnState: func(c net.Conn, s http.ConnState) {
			log.Printf("➡️  %s → %v", c.RemoteAddr(), s)
		},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server ouvindo em http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
```

- **O que observar**  
  - Para cada transição de estado (`StateNew`, `StateActive`, `StateIdle`, `StateClosed`) você verá um log.  
  - A função anônima em `HandlerFunc` demonstra a simplicidade do modelo: só precisa escrever no `ResponseWriter`.  
  - Ajuste `ReadTimeout`/`IdleTimeout` e veja como o servidor fecha conexões inativas.

---

### Desafios para praticar 🚀

1. **Middleware de log**  
   Crie um *middleware* que meça o tempo de cada requisição, registrando método, rota, código de status e duração.

2. **Cancelamento via `context`**  
   Modifique o handler para simular uma tarefa longa (ex.: `time.Sleep(3 * time.Second)`) e cancele se o cliente fechar a conexão antes.

3. **HTTP/2 + TLS**  
   Gere um certificado autoassinado, ative `http.Server.TLSConfig` e compare o comportamento de *keep‑alive* em HTTP/1.1 vs HTTP/2.

4. **Pool de conexões do cliente**  
   Escreva um *benchmark* com `http.Client{Transport:&http.Transport{MaxIdleConns: 100}}` e veja como o reuso reduz latência.

5. **Expondo métricas Prometheus**  
   Integre `promhttp.Handler()` para expor métricas e inspecione `go_http_requests_total`, `go_goroutines`, etc.