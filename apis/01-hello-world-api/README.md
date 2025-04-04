# Hello World API

Este é o primeiro projeto da série de APIs em Go. Uma API simples que retorna "Hello, World!" quando acessada.

## Objetivos de Aprendizado

- Configurar um projeto Go básico
- Criar um servidor HTTP simples
- Implementar um endpoint GET
- Estruturar o projeto de forma organizada
- Implementar testes básicos
- Documentar a API

## Tecnologias Utilizadas

- Go 1.21+
- [Gin Web Framework](https://gin-gonic.com/)
- Go Modules
- Testing package

## Estrutura do Projeto

```
01-hello-world-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── handlers/
│       └── hello.go
├── tests/
│   └── hello_test.go
├── go.mod
├── go.sum
└── README.md
```

## Como Executar

1. Clone o repositório
2. Entre no diretório do projeto:
   ```bash
   cd apis/01-hello-world-api
   ```
3. Instale as dependências:
   ```bash
   go mod tidy
   ```
4. Execute o projeto:
   ```bash
   go run cmd/api/main.go
   ```
5. Acesse a API:
   ```bash
   curl http://localhost:8080/hello
   ```

## Endpoints

### GET /hello

Retorna uma mensagem de "Hello, World!"

**Resposta de Sucesso:**
```json
{
    "message": "Hello, World!"
}
```

## Testes

Para executar os testes:
```bash
go test ./tests/...
```

## Próximos Passos

1. Adicionar mais endpoints
2. Implementar logging
3. Adicionar middleware para CORS
4. Implementar rate limiting
5. Adicionar documentação com Swagger 