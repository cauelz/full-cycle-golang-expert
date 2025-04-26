package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo")
	})

	error := http.ListenAndServe(":8080", nil)

	if error != nil {
		panic(error)
	}
}