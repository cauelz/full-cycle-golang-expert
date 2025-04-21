# Middlewares em Go - Exemplo Prático

Este exemplo demonstra o uso de middlewares em Go através de uma API de tarefas. O exemplo implementa vários tipos de middlewares e mostra como encadeá-los de forma eficiente.

## Funcionalidades

1. **API de Tarefas**
   - Listar tarefas (GET /tasks)
   - Criar tarefa (POST /tasks)
   - Armazenamento em memória thread-safe

2. **Middlewares Implementados**
   - Logging (registro de requisições)
   - Recovery (recuperação de pânicos)
   - Timeout (limite de tempo para requisições)
   - CORS (Cross-Origin Resource Sharing)

3. **Características**
   - Encadeamento de middlewares
   - Wrapper de ResponseWriter
   - Graceful shutdown
   - Tratamento de sinais

## Como Executar

1. Clone o repositório e entre no diretório:
   ```bash
   cd exemplos/04-middlewares
   ```

2. Execute o programa:
   ```bash
   go run main.go
   ```

## Testando a API

1. **Listar Tarefas**
   ```bash
   curl http://localhost:8080/tasks
   ```

2. **Criar Tarefa**
   ```bash
   curl -X POST http://localhost:8080/tasks \
     -H "Content-Type: application/json" \
     -d '{"title": "Estudar Go", "done": false}'
   ```

3. **Testar CORS**
   ```bash
   curl -X OPTIONS http://localhost:8080/tasks \
     -H "Origin: http://example.com" \
     -v
   ```

## Exemplo de Saída

```
2024/01/01 12:00:00 Servidor iniciado em http://localhost:8080
2024/01/01 12:00:05 method=GET path=/tasks status=200 duration=1.234ms
2024/01/01 12:00:10 method=POST path=/tasks status=201 duration=2.345ms
```

## Estrutura do Código

1. **Middlewares**
   - `LoggingMiddleware`: registra informações sobre requisições
   - `RecoveryMiddleware`: recupera de pânicos
   - `TimeoutMiddleware`: adiciona timeout para requisições
   - `CORSMiddleware`: configura headers CORS

2. **Componentes Principais**
   - `TaskStore`: armazenamento thread-safe
   - `TaskHandler`: gerenciamento de rotas
   - `Chain`: função de encadeamento
   - `responseWriter`: wrapper para captura de status

3. **Boas Práticas**
   - Uso de interfaces
   - Encapsulamento
   - Tratamento de erros
   - Graceful shutdown

## Características dos Middlewares

1. **LoggingMiddleware**
   - Registra método, path, status e duração
   - Usa wrapper de ResponseWriter
   - Logging estruturado

2. **RecoveryMiddleware**
   - Recupera de pânicos
   - Registra stack trace
   - Retorna 500 Internal Server Error

3. **TimeoutMiddleware**
   - Configura timeout por requisição
   - Usa context para cancelamento
   - Retorna 504 Gateway Timeout

4. **CORSMiddleware**
   - Configura headers CORS
   - Trata preflight requests
   - Permite customização de origens

## Próximos Passos

1. Adicionar mais middlewares:
   - Autenticação
   - Rate limiting
   - Compressão
   - Cache

2. Melhorias:
   - Configuração via arquivo
   - Métricas detalhadas
   - Testes unitários
   - Documentação com Swagger 