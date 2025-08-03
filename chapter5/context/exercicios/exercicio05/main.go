// Exemplo de uso de context.WithCancel em Go
//
// Você pode aguardar o cancelamento de um contexto de duas formas principais:
//
// 1. Usando apenas <-ctx.Done(), que bloqueia até o contexto ser cancelado (semelhante ao await de uma Promise no JavaScript).
// 2. Usando select, útil quando você quer aguardar múltiplos canais ao mesmo tempo.
//
// O comportamento de <-ctx.Done() lembra o async/await do JavaScript, pois você "espera" até que algo aconteça (no caso, o cancelamento do contexto).
//
// Exemplos abaixo:

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()

	// Exemplo 1: Esperando apenas o cancelamento do contexto (simples, como await)
	fmt.Println("Aguardando cancelamento do contexto (exemplo 1)...")
	<-ctx.Done()
	fmt.Println("Contexto finalizado pelo cancelamento manual (exemplo 1).")

	// Exemplo 2: Usando select para aguardar múltiplos canais
	ctx2, cancel2 := context.WithCancel(context.Background())
	outroCanal := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		outroCanal <- "Mensagem recebida!"
	}()
	go func() {
		time.Sleep(time.Second * 4)
		cancel2()
	}()

	fmt.Println("Aguardando contexto ou mensagem (exemplo 2)...")
	select {
	case <-ctx2.Done():
		fmt.Println("Contexto cancelado (exemplo 2).")
	case msg := <-outroCanal:
		fmt.Println("Recebi do canal (exemplo 2):", msg)
	}
}
