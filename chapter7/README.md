# Desafio: Cotação do Dólar em Go

Olá dev, tudo bem?

Neste desafio, vamos aplicar o que aprendemos sobre:
- Webserver HTTP
- Contextos
- Banco de dados
- Manipulação de arquivos com Go

## Objetivo

Você precisará entregar **dois sistemas em Go**:
- `client.go`
- `server.go`

## Requisitos

### 1. `client.go`
- Deve realizar uma requisição HTTP ao `server.go` solicitando a cotação do dólar.
- Precisa receber do `server.go` apenas o valor atual do câmbio (campo `bid` do JSON).
- Utilizando o package `context`, terá um timeout máximo de **300ms** para receber o resultado do `server.go`.
- Deve salvar a cotação atual em um arquivo `cotacao.txt` no formato:
  
  ```
  Dólar: {valor}
  ```

### 2. `server.go`
- Deve consumir a API de câmbio Dólar/Real no endereço: [https://economia.awesomeapi.com.br/json/last/USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL)
- Deve retornar o resultado para o cliente no formato JSON.
- Usando o package `context`, deve registrar no banco de dados SQLite cada cotação recebida.
  - O timeout máximo para chamar a API de cotação do dólar deve ser de **200ms**.
  - O timeout máximo para persistir os dados no banco deve ser de **10ms**.
- O endpoint necessário será `/cotacao` e a porta do servidor HTTP será **8080**.

### 3. Contextos e Logs
- Os **3 contextos** (client, chamada da API e persistência no banco) devem retornar erro nos logs caso o tempo de execução seja insuficiente.

## Entrega

Ao finalizar, envie o link do repositório para correção.