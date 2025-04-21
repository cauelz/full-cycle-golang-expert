# Exemplo de Servidor HTTP Básico em Go

Este é um exemplo simples de um servidor HTTP em Go que implementa um serviço de tarefas (TODO) com armazenamento em memória.

## Funcionalidades

- Servidor HTTP com configurações de timeout
- Graceful shutdown
- Manipulação de JSON
- Rotas para listar e criar tarefas
- Armazenamento em memória

## Como Executar

1. Entre no diretório do exemplo:
   ```bash
   cd exemplos/01-servidor-basico
   ```

2. Execute o servidor:
   ```bash
   go run main.go
   ```

3. O servidor estará disponível em `http://localhost:8080`

## Endpoints

### GET /tasks
Lista todas as tarefas.

Exemplo de resposta:
```json
[
  {
    "id": 1,
    "title": "Aprender Go",
    "done": false,
    "created_at": "2024-03-20T10:00:00Z"
  }
]
```

### POST /tasks
Cria uma nova tarefa.

Exemplo de requisição:
```json
{
  "title": "Aprender Go"
}
```

Exemplo de resposta:
```json
{
  "id": 1,
  "title": "Aprender Go",
  "done": false,
  "created_at": "2024-03-20T10:00:00Z"
}
```

## Testando com cURL

1. Listar tarefas:
   ```bash
   curl http://localhost:8080/tasks
   ```

2. Criar tarefa:
   ```bash
   curl -X POST http://localhost:8080/tasks \
        -H "Content-Type: application/json" \
        -d '{"title": "Aprender Go"}'
   ```

## Características do Código

1. **Configurações de Timeout**
   - ReadTimeout: 15 segundos
   - WriteTimeout: 15 segundos
   - IdleTimeout: 60 segundos

2. **Graceful Shutdown**
   - Captura sinais SIGINT e SIGTERM
   - Aguarda requisições em andamento
   - Timeout de 5 segundos para encerramento

3. **Estrutura do Código**
   - Separação em tipos e métodos
   - Uso de interfaces do Go
   - Tratamento de erros apropriado

## Próximos Passos

1. Adicionar mais operações (UPDATE, DELETE)
2. Implementar persistência em banco de dados
3. Adicionar validações
4. Implementar autenticação
5. Adicionar testes 