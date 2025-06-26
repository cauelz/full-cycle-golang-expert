package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second * 5))
	defer cancel()

	select {
	case <- ctx.Done():
		fmt.Println("Execução finalizada pelo deadline.")
	case <- time.After(time.Second * 10):
		fmt.Println("Execução finalizada pelo temporizador.")
	}
}