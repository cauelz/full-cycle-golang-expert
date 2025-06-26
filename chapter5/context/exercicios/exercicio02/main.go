package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	select {
	case <- ctx.Done():
		fmt.Println("Processamento Finalizado pelo cancelamento do contexto")

	case <- time.After(time.Second * 3):
		fmt.Println("Cancelamento apos 3 segundos")
	}
}