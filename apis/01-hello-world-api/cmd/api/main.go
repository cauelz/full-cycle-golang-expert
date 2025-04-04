package main

import (
	"log"

	"hello-world/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Criar uma nova instância do Gin
	r := gin.Default()

	// Configurar as rotas
	r.GET("/hello", handlers.HelloHandler)

	// Iniciar o servidor na porta 8080
	log.Println("Servidor iniciando na porta 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
} 