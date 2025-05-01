# Trabalhando com o pacote Context em Go

O pacote `context` é um dos mais importantes na linguagem Go, pois permite criar e propagar informações de contexto entre goroutines, controlando o ciclo de vida de processos, especialmente em aplicações concorrentes e servidores web.

A partir de um contexto, podemos:

- Controlar e limitar o tempo de execução de um processo (timeout ou deadline).
- Cancelar operações em andamento de forma coordenada, poupando recursos do sistema.
- Propagar sinais de cancelamento entre diferentes partes do código.
- Compartilhar valores e metadados entre funções e goroutines de forma segura.

## Principais funções e conceitos

- **context.Background()**: Retorna um contexto vazio, geralmente usado como contexto raiz.
- **context.TODO()**: Usado quando ainda não se sabe qual contexto usar.
- **context.WithCancel(parent)**: Cria um novo contexto derivado do contexto pai, que pode ser cancelado manualmente.
- **context.WithTimeout(parent, duration)**: Cria um contexto que será automaticamente cancelado após um determinado tempo.
- **context.WithDeadline(parent, time)**: Cria um contexto que será cancelado em um horário específico.
- **context.WithValue(parent, key, value)**: Permite armazenar e recuperar valores associados ao contexto.

## Exemplo prático

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

select {
case <-time.After(3 * time.Second):
    fmt.Println("Processo concluído")
case <-ctx.Done():
    fmt.Println("Timeout ou cancelamento:", ctx.Err())
}
```

Neste exemplo, o contexto será cancelado após 2 segundos, interrompendo o processo caso ele demore mais do que isso.

## Boas práticas

- Sempre propague o contexto como o primeiro parâmetro das funções: `func MinhaFuncao(ctx context.Context, ...)`.
- Evite usar `context.WithValue` para passar dados de negócio; use apenas para informações de controle, como IDs de requisição.
- Sempre chame a função de cancelamento (`cancel()`) para liberar recursos.