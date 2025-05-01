# Exemplo: Controle de Timeout com Context e Select em Go

Este exemplo demonstra como utilizar o pacote `context` do Go para controlar o tempo de execução de uma operação, aplicando um timeout, e como o comando `select` pode ser usado para lidar com múltiplos canais de forma concorrente.

## Objetivo

Simular uma reserva de hotel que pode ser cancelada automaticamente caso ultrapasse um tempo limite (timeout). O exemplo mostra como:
- Criar um contexto com timeout
- Cancelar operações longas automaticamente
- Utilizar o `select` para aguardar múltiplos eventos concorrentes

## Código-fonte

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <- ctx.Done():
		fmt.Println("Hotel Booking cancelled. Timeout reached")
		return
	case <- time.After(1 * time.Second):
		fmt.Println("Hotel booked")
		return
	}
}
```

## Explicação dos Conceitos

### 1. Context e Timeout
O pacote `context` é fundamental para controlar deadlines, cancelamentos e propagação de sinais entre goroutines. No exemplo:
- `context.Background()` cria um contexto "vazio", raiz.
- `context.WithTimeout(ctx, time.Second*3)` cria um novo contexto derivado, que será automaticamente cancelado após 3 segundos.
- `cancel()` é chamado com `defer` para garantir que recursos sejam liberados ao final da função, mesmo que o timeout não seja atingido.

### 2. Função bookHotel e o uso do select
A função `bookHotel` recebe o contexto e utiliza o comando `select`:

```go
select {
case <- ctx.Done():
	fmt.Println("Hotel Booking cancelled. Timeout reached")
	return
case <- time.After(1 * time.Second):
	fmt.Println("Hotel booked")
	return
}
```

#### O que é o select?
O `select` em Go permite que você aguarde múltiplas operações de canal ao mesmo tempo. Ele escolhe aleatoriamente um dos cases prontos para execução. Se mais de um canal estiver pronto, um deles é escolhido aleatoriamente. Se nenhum estiver pronto, o select bloqueia até que algum esteja.

#### Como funciona neste exemplo?
- `case <- ctx.Done()`: Este canal é fechado quando o contexto é cancelado (por timeout ou manualmente). Se o timeout de 3 segundos for atingido antes da reserva ser concluída, esta branch será executada.
- `case <- time.After(1 * time.Second)`: Cria um canal que "dispara" após 1 segundo. Se a reserva for concluída antes do timeout, esta branch será executada.

### 3. Fluxo do Programa
- O programa inicia o contexto com timeout de 3 segundos.
- Chama `bookHotel`, que espera 1 segundo simulando a reserva.
- Como 1 segundo < 3 segundos, a mensagem será `Hotel booked`.
- Se aumentarmos o tempo de reserva para mais de 3 segundos, veremos `Hotel Booking cancelled. Timeout reached`.

## Testando o Timeout
Para testar o timeout, altere a linha:
```go
case <- time.After(1 * time.Second):
```
para, por exemplo:
```go
case <- time.After(4 * time.Second):
```
Agora, o timeout será atingido antes da reserva ser concluída, e a mensagem de cancelamento será exibida.

## Resumo
- Use `context` para controlar cancelamentos e timeouts em operações concorrentes.
- Use `select` para aguardar múltiplos canais e tratar diferentes cenários de execução.
- Sempre libere recursos com `defer cancel()` ao usar contextos com timeout ou cancelamento.

---

**Referências:**
- [Documentação oficial do context](https://pkg.go.dev/context)
- [Tour of Go: Select](https://go.dev/tour/concurrency/5) 