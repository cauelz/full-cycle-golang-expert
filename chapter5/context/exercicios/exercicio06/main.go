package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "senha", "1234")

	done := make(chan struct{})

	go func(ctx context.Context) {
		time.Sleep(time.Second * 5)
		fmt.Println(ctx.Value("senha"))
		close(done)
	}(ctx)

	select {
	case <- done:
		fmt.Println("Canal finalizado")
	}
}
