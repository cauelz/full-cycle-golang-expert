# Guia do Exemplo 10: Trabalhando com Timeout em integrações HTTP (Go)

## Objetivo
Este exemplo demonstra como utilizar timeouts em requisições HTTP feitas por um cliente em Go. O timeout é fundamental para evitar que sua aplicação fique travada aguardando respostas de serviços externos que podem estar lentos ou indisponíveis.

## O que o código faz?
O código possui duas funções principais:

- `callGet()`: Realiza uma requisição GET para `https://google.com` utilizando um cliente HTTP com timeout extremamente baixo (1 microssegundo), forçando o disparo do timeout.
- `callPost()`: Realiza uma requisição POST para `https://google.com` enviando um JSON simples, mas sem configurar timeout explícito (usa o padrão do pacote `http`).

## Explicação do Timeout
No Go, o timeout pode ser configurado diretamente no struct `http.Client`. Se o tempo de resposta do servidor for maior que o timeout definido, a requisição é abortada e um erro é retornado.

Exemplo do código:
```go
c := http.Client{
    Timeout: time.Microsecond, // Timeout extremamente baixo
}
resp, err := c.Get("https://google.com")
```

Neste caso, o timeout é tão baixo que a requisição sempre irá falhar, servindo para demonstrar o tratamento de erro.

## Como executar o exemplo
1. Certifique-se de ter o Go instalado.
2. No terminal, navegue até a pasta deste exemplo:
   ```sh
   cd chapter5/http/exemplos/10-http-client-timeout
   ```
3. Execute o programa:
   ```sh
   go run main.go
   ```

Você verá um panic na função `callGet()` devido ao timeout. Se comentar a chamada de `callGet()` e deixar apenas `callPost()`, verá a resposta do Google (provavelmente um HTML de erro, pois o endpoint não aceita POST com JSON).

## Pontos de Atenção
- Sempre defina timeouts em integrações externas para evitar travamentos.
- O valor do timeout deve ser razoável para o contexto da sua aplicação.
- Trate erros de timeout de forma adequada para melhorar a resiliência do seu sistema.

## Sugestão de Experimentos
- Altere o valor do timeout para 1 ou 2 segundos e veja o comportamento.
- Troque a URL para um endpoint que aceite POST e observe a resposta.

---

Este exemplo é didático e serve para reforçar a importância do uso de timeout em integrações HTTP no Go. 