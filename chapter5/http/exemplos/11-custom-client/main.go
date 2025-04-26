package main

import (
	"io"
	"net/http"
)

func main() {

	c := http.Client{}

	req, error := http.NewRequest("GET", "http://google.com", nil)

	if error != nil {
		panic(error)
	}

	req.Header.Set("Accept", "application/json")

	resp, error := c.Do(req)

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)

	if error != nil {
		panic(error)
	}

	println(string(body))
}