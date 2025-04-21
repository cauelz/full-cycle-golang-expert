# Cliente HTTP em Go

Este exemplo demonstra a implementação de um cliente HTTP em Go com recursos avançados e boas práticas para aplicações em produção.

## Características

- **Cliente HTTP Otimizado**: Configurações personalizadas de timeout, conexões e pool
- **Pool de Buffers**: Reutilização de buffers para reduzir a pressão no GC
- **Tratamento de Contexto**: Suporte a cancelamento e timeouts
- **Retry Automático**: Implementação de retry com backoff exponencial
- **Tratamento de Erros**: Mensagens de erro detalhadas e tratamento adequado
- **Upload de Arquivos**: Suporte a upload de arquivos via multipart/form-data
- **Query Parameters**: Suporte a parâmetros de query string
- **Headers Personalizados**: Configuração flexível de headers HTTP

## Pré-requisitos

- Go 1.22 ou superior
- Servidor de exemplo rodando (use o exemplo `03-performance` ou `02-api-rest`)

## Executando o Cliente

1. Navegue até o diretório do exemplo:
   ```bash
   cd chapter5/http/exemplos/06-http-client
   ```

2. Execute o cliente:
   ```bash
   go run main.go
   ```

## Funcionalidades Implementadas

### 1. Buscar Produtos
```go
products, err := client.GetProducts(ctx)
```
- Faz uma requisição GET para `/products`
- Retorna uma lista de produtos
- Suporta cancelamento via contexto

### 2. Criar Produto
```go
product := Product{
    Name:        "New Product",
    Description: "A new product",
    Price:       99.99,
}
err := client.CreateProduct(ctx, product)
```
- Faz uma requisição POST para `/products`
- Envia dados JSON no body
- Usa pool de buffers para otimização

### 3. Buscar com Filtros
```go
query := url.Values{}
query.Set("minPrice", "50")
query.Set("maxPrice", "100")
products, err := client.SearchProducts(ctx, query)
```
- Suporta query parameters
- Codifica parâmetros automaticamente
- Retorna produtos filtrados

### 4. Upload de Arquivo
```go
err := client.UploadFile(ctx, "example.txt")
```
- Suporta upload de arquivos via multipart/form-data
- Gerencia recursos corretamente (fechamento de arquivos)
- Configura headers apropriados

## Características do Código

### Configurações do Cliente

```go
client := &http.Client{
    Timeout: time.Second * 30,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
        MaxConnsPerHost:     100,
        IdleConnTimeout:     90 * time.Second,
        TLSHandshakeTimeout: 10 * time.Second,
    },
}
```

- Timeout global de 30 segundos
- Pool de conexões configurado
- Timeouts específicos para diferentes fases da conexão

### Pool de Buffers

```go
bufferPool: &sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}
```

- Reutilização de buffers
- Reduz alocações de memória
- Thread-safe

### Tratamento de Erros

- Mensagens de erro detalhadas
- Informações de status code e body em erros
- Fechamento adequado de recursos

## Testando o Cliente

1. Primeiro, inicie o servidor de exemplo:
   ```bash
   cd ../03-performance
   go run main.go
   ```

2. Em outro terminal, execute o cliente:
   ```bash
   cd ../06-http-client
   go run main.go
   ```

## Próximos Passos

1. Adicionar autenticação
2. Implementar circuit breaker
3. Adicionar métricas de requests
4. Implementar cache de respostas
5. Adicionar logging estruturado
6. Implementar rate limiting
7. Adicionar tracing distribuído
8. Implementar testes de integração
9. Adicionar suporte a webhooks
10. Implementar retry configurável 