# Servidor HTTP Otimizado para Performance em Go

Este exemplo demonstra uma implementação de servidor HTTP de alta performance em Go, incorporando várias técnicas de otimização e melhores práticas para aplicações em nível de produção.

## Funcionalidades

- **Cache em Memória**: Implementa um cache thread-safe com TTL e limpeza automática
- **Compressão de Resposta**: Middleware de compressão Gzip para reduzir o tamanho das respostas
- **Métricas de Performance**: Monitoramento em tempo real de contagem de requisições, tempos de resposta e hits no cache
- **Otimização de Memória**: Usa sync.Pool para reutilização de buffers reduzindo a pressão no GC
- **Desligamento Gracioso**: Tratamento adequado de desligamento com cancelamento de contexto
- **Suporte a Profiling**: Endpoints pprof integrados para análise de performance
- **Tratamento de Requisições Concorrentes**: Gerenciamento eficiente de goroutines
- **Logs Estruturados**: Logs detalhados de requisições e erros

## Pré-requisitos

- Go 1.22 ou superior

## Executando o Servidor

1. Navegue até o diretório do exemplo:
   ```bash
   cd chapter5/http/exemplos/03-performance
   ```

2. Execute o servidor:
   ```bash
   go run main.go
   ```

O servidor iniciará na porta 8080.

## Endpoints Disponíveis

### 1. Obter Produto
```
GET /products/{id}
```
Retorna um produto por ID. As respostas são cacheadas por 30 segundos.

Exemplo:
```bash
curl http://localhost:8080/products/1
```

### 2. Obter Métricas
```
GET /metrics
```
Retorna métricas atuais do servidor incluindo:
- Total de requisições
- Tempo médio de resposta
- Taxa de acerto do cache
- Conexões ativas

Exemplo:
```bash
curl http://localhost:8080/metrics
```

### 3. Endpoints de Profiling
Disponíveis em `/debug/pprof/`:
- Perfil de memória: `/debug/pprof/heap`
- Perfil de CPU: `/debug/pprof/profile`
- Perfil de goroutines: `/debug/pprof/goroutine`

Exemplo de coleta de um perfil de CPU de 30 segundos:
```bash
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
```

## Explicação das Funcionalidades de Performance

### 1. Sistema de Cache
- Cache em memória thread-safe com expiração
- Limpeza automática de itens expirados
- Tempo de busca O(1) para respostas cacheadas

### 2. Compressão de Resposta
- Compressão Gzip automática para respostas
- Decisões de compressão baseadas no Content-Type
- Respeita o cabeçalho Accept-Encoding do cliente

### 3. Pool de Buffers
- Reutiliza buffers de resposta para reduzir alocações de memória
- Minimiza a pressão no garbage collector
- Implementação thread-safe usando sync.Pool

### 4. Coleta de Métricas
- Contadores atômicos de baixo overhead
- Monitoramento de performance em tempo real
- Rastreamento da efetividade do cache

### 5. Desligamento Gracioso
- Aguarda a conclusão das requisições em andamento
- Limpeza adequada de recursos
- Tratamento de sinais (SIGINT/SIGTERM)

## Teste de Carga

Você pode usar ferramentas como `hey` ou `wrk` para testar a carga do servidor:

```bash
# Instalar hey
go install github.com/rakyll/hey@latest

# Executar um teste de carga (100 usuários concorrentes, 10000 requisições)
hey -n 10000 -c 100 http://localhost:8080/products/1
```

## Monitorando Performance

1. Visualizar métricas em tempo real:
```bash
watch -n 1 'curl -s http://localhost:8080/metrics'
```

2. Gerar e analisar perfil de CPU:
```bash
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile
```

3. Análise do perfil de memória:
```bash
go tool pprof -http=:8081 http://localhost:8080/debug/pprof/heap
```

## Próximos Passos

1. Adicionar cache distribuído (ex: Redis)
2. Implementar limitação de taxa
3. Adicionar rastreamento OpenTelemetry
4. Configurar exportação de métricas para Prometheus
5. Adicionar circuit breaker para dependências externas
6. Implementar middleware de timeout para requisições
7. Adicionar endpoints de verificação de saúde
8. Implementar cache de respostas baseado em cabeçalhos
9. Adicionar middleware de validação de requisições
10. Implementar mecanismos de retry para operações falhas 