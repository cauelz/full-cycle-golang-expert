# API REST em Go com SQLite

Este exemplo demonstra uma implementação completa de uma API REST em Go usando SQLite como banco de dados. O exemplo inclui:

- Arquitetura em camadas (handler, service, database)
- CRUD completo de usuários
- Tratamento de erros
- Validações
- Middleware de logging
- Graceful shutdown

## Pré-requisitos

1. Go 1.22 ou superior
2. SQLite3

## Instalação

1. Clone o repositório
2. Instale as dependências:
   ```bash
   go mod init api-rest
   go get github.com/mattn/go-sqlite3
   ```

## Estrutura do Projeto

```
.
├── main.go     # Código principal
├── users.db    # Banco de dados SQLite (criado automaticamente)
└── README.md   # Este arquivo
```

## Como Executar

1. Entre no diretório do exemplo:
   ```bash
   cd exemplos/02-api-rest
   ```

2. Execute o servidor:
   ```bash
   go run main.go
   ```

3. O servidor estará disponível em `http://localhost:8080`

## Endpoints

### GET /api/users
Lista todos os usuários.

Exemplo de resposta:
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao@exemplo.com",
    "created_at": "2024-03-20T10:00:00Z"
  }
]
```

### GET /api/users/{id}
Retorna um usuário específico.

Exemplo de resposta:
```json
{
  "id": 1,
  "name": "João Silva",
  "email": "joao@exemplo.com",
  "created_at": "2024-03-20T10:00:00Z"
}
```

### POST /api/users
Cria um novo usuário.

Exemplo de requisição:
```json
{
  "name": "João Silva",
  "email": "joao@exemplo.com"
}
```

### PUT /api/users/{id}
Atualiza um usuário existente.

Exemplo de requisição:
```json
{
  "name": "João Silva Atualizado",
  "email": "joao.novo@exemplo.com"
}
```

### DELETE /api/users/{id}
Remove um usuário.

## Testando com cURL

1. Listar usuários:
   ```bash
   curl http://localhost:8080/api/users
   ```

2. Criar usuário:
   ```bash
   curl -X POST http://localhost:8080/api/users \
        -H "Content-Type: application/json" \
        -d '{"name": "João Silva", "email": "joao@exemplo.com"}'
   ```

3. Buscar usuário:
   ```bash
   curl http://localhost:8080/api/users/1
   ```

4. Atualizar usuário:
   ```bash
   curl -X PUT http://localhost:8080/api/users/1 \
        -H "Content-Type: application/json" \
        -d '{"name": "João Silva Atualizado", "email": "joao.novo@exemplo.com"}'
   ```

5. Remover usuário:
   ```bash
   curl -X DELETE http://localhost:8080/api/users/1
   ```

## Características do Código

1. **Arquitetura em Camadas**
   - Handlers: Lida com HTTP
   - Service: Lógica de negócio
   - Database: Persistência

2. **Tratamento de Erros**
   - Validação de entrada
   - Erros de negócio
   - Erros de banco de dados

3. **Middleware**
   - Logging de requisições
   - Fácil de adicionar mais middlewares

4. **Graceful Shutdown**
   - Espera requisições em andamento
   - Fecha conexões corretamente

## Próximos Passos

1. Adicionar autenticação
2. Implementar cache
3. Adicionar testes
4. Documentar API com Swagger
5. Adicionar métricas e tracing 