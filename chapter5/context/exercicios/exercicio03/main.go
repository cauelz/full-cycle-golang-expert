package main

import (
	"context"
	"fmt"
	"time"
)


func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	select {
	case <- ctx.Done():
		fmt.Println("Contexto Finalizado pelo timeout")
	case <- time.After(time.Second * 2 ):
		fmt.Println("Execução finalizada pelo temporizador")
	}
}
