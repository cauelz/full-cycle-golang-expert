package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Iniciado.")

	defer log.Println("Request Finalizado.")

	select{
	case <- time.After(5 * time.Second):
		log.Println("Request processado com sucesso.")
		w.Write([]byte("Request processado com sucesso."))
	case <- ctx.Done():
		log.Println("Request cancelado pelo cliente!")
	}
}