package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if error != nil {
		panic(error)
	}

	resp, error := http.DefaultClient.Do(req)

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}