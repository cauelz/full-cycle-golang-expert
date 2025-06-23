package main

import "net/http"

func main() {


}

func NewServer() {

	error := http.ListenAndServe(":8080", nil)

	if error != nil {
		panic(error)
	}
}