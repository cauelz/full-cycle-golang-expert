# Exemplo: Client HTTP com Contexto em Go

Este exemplo demonstra como criar um cliente HTTP em Go que faz uma requisição GET para um servidor local (`http://localhost:8080`) utilizando o pacote `context` para controlar o tempo limite (timeout) da requisição.

## Descrição do Código

O código do arquivo `main.go` faz o seguinte:

1. **Criação do Contexto com Timeout:**
   - Utiliza `context.WithTimeout` para criar um contexto com tempo limite de 10 segundos. Isso significa que, se a requisição não for concluída em até 10 segundos, ela será automaticamente cancelada.

2. **Criação da Requisição HTTP:**
   - Cria uma requisição GET usando `http.NewRequestWithContext`, associando o contexto criado anteriormente à requisição.

3. **Envio da Requisição:**
   - Usa `http.DefaultClient.Do` para enviar a requisição ao servidor.
   - Se ocorrer algum erro (por exemplo, se o servidor não estiver rodando ou se o tempo limite for atingido), o programa irá interromper a execução com `panic`.

4. **Leitura da Resposta:**
   - O corpo da resposta (`resp.Body`) é impresso diretamente no terminal usando `io.Copy(os.Stdout, resp.Body)`.

## Como Executar

1. Certifique-se de que há um servidor HTTP rodando em `http://localhost:8080`.
2. Execute o programa:

```bash
cd chapter5/context/exemplos/02-server-client
go run client/main.go
```

Se o servidor responder dentro de 10 segundos, a resposta será exibida no terminal. Caso contrário, o programa será encerrado devido ao timeout.

## Pontos Importantes

- O uso de `context.WithTimeout` é fundamental para evitar que o cliente fique esperando indefinidamente por uma resposta do servidor.
- O contexto pode ser usado para cancelar a requisição manualmente ou por outros motivos além do timeout.

---

Este exemplo é útil para entender como implementar controle de tempo e cancelamento em requisições HTTP usando Go. 