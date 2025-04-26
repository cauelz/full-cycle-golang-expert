package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	callGet()
	callPost()

}

func callGet() {
	c := http.Client{
		Timeout: time.Microsecond,
	}

	resp, error := c.Get("https://google.com")

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close();

	body, error := io.ReadAll(resp.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))
}

func callPost() {

	json := bytes.NewBuffer([]byte(`{"name": "Caue Zaratin"}`))

	resp, error := http.Post("https://google.com", "application/json", json)

	if error != nil {
		panic(error)
	}

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}