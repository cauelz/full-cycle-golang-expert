# Exemplo: Cancelamento de Request com Contexto no Go

Este exemplo demonstra como utilizar o contexto (`context.Context`) em um servidor HTTP em Go para lidar com o cancelamento de requisições pelo cliente.

## Funcionamento

- O servidor HTTP escuta na porta `8080` e responde a todas as requisições na raiz (`/`).
- Ao receber uma requisição, o handler:
  1. Obtém o contexto da requisição (`r.Context()`).
  2. Loga que o request foi iniciado.
  3. Utiliza um `select` para aguardar:
     - **5 segundos** (simulando um processamento demorado), ou
     - O cancelamento do contexto (por exemplo, se o cliente fechar a conexão antes do tempo).
  4. Se o processamento terminar antes do cancelamento, responde ao cliente com sucesso.
  5. Se o contexto for cancelado antes, loga que o request foi cancelado pelo cliente.
  6. Ao final, loga que o request foi finalizado (usando `defer`).

## Código Principal
```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    log.Println("Request Iniciado.")
    defer log.Println("Request Finalizado.")
    select {
    case <-time.After(5 * time.Second):
        log.Println("Request processado com sucesso.")
        w.Write([]byte("Request processado com sucesso."))
    case <-ctx.Done():
        log.Println("Request cancelado pelo cliente!")
    }
}
```

## Como testar
1. Execute o servidor:
   ```sh
   go run main.go
   ```
2. Em outro terminal, faça uma requisição:
   ```sh
   curl localhost:8080
   ```
   - Se aguardar 5 segundos, verá a resposta "Request processado com sucesso.".
   - Se cancelar a requisição antes (Ctrl+C no curl), o servidor logará "Request cancelado pelo cliente!".

## Objetivo
Este exemplo mostra como o contexto pode ser usado para tornar o servidor mais eficiente, liberando recursos imediatamente quando o cliente desiste da requisição. 