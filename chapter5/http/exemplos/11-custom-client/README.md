# Exemplo: Cliente HTTP Customizado em Go

Este exemplo demonstra como criar um cliente HTTP customizado utilizando as bibliotecas padrão do Go (`net/http`). O código faz uma requisição GET para o site do Google, define um header customizado e lê a resposta.

## Código de Exemplo

```go
package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}

	req, error := http.NewRequest("GET", "http://google.com", nil)
	if error != nil {
		panic(error)
	}

	req.Header.Set("Accept", "application/json")

	resp, error := c.Do(req)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}

	println(string(body))
}
```

## Explicação do Exemplo

1. **Criação do Cliente HTTP**:
   - `c := http.Client{}` cria uma instância do cliente HTTP. Você pode customizar esse cliente, por exemplo, definindo timeouts, proxies, ou transportadores customizados.

2. **Criação da Requisição**:
   - `http.NewRequest("GET", "http://google.com", nil)` cria uma nova requisição HTTP do tipo GET. O terceiro parâmetro pode ser usado para enviar um corpo (body) na requisição, útil para métodos como POST ou PUT.

3. **Customização de Headers**:
   - `req.Header.Set("Accept", "application/json")` define um header customizado na requisição. Você pode adicionar qualquer header necessário para sua integração.

4. **Envio da Requisição**:
   - `c.Do(req)` envia a requisição usando o cliente HTTP criado.

5. **Leitura da Resposta**:
   - `io.ReadAll(resp.Body)` lê todo o corpo da resposta. Não se esqueça de fechar o body após a leitura (`defer resp.Body.Close()`).

## Como criar integrações customizadas com http.Client e http.Request

A biblioteca `net/http` do Go é bastante flexível e permite customizar tanto o cliente quanto as requisições. Veja algumas dicas:

### Customizando o http.Client

Você pode definir opções como timeout, proxy, e transporte customizado:

```go
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        // Configurações avançadas, como proxy, TLS, etc.
    },
}
```

### Customizando a http.Request

- Adicione headers customizados conforme necessário:
  ```go
  req.Header.Set("Authorization", "Bearer <token>")
  req.Header.Set("Content-Type", "application/json")
  ```
- Para métodos como POST ou PUT, envie um body:
  ```go
  body := strings.NewReader(`{"key":"value"}`)
  req, err := http.NewRequest("POST", url, body)
  ```

### Exemplos de uso avançado

- **Autenticação**: Adicione tokens ou cookies nos headers.
- **Timeouts**: Defina timeouts no cliente para evitar requisições travadas.
- **Transporte customizado**: Implemente um transporte customizado para logging, métricas ou manipulação de baixo nível das requisições.

## Referências
- [Documentação oficial do net/http](https://pkg.go.dev/net/http)
- [Como customizar o http.Client](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)

---

Sinta-se à vontade para modificar este exemplo para atender às necessidades da sua integração! 