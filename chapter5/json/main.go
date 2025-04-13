package main

import (
	"fmt"
	"os"
)

func main() {

	// Escolha qual exemplo executar.

	if len(os.Args) < 2 {
		fmt.Println("Escolha o exemplo para ser executado:")
		fmt.Println("1 - Cria uma instancia de Pessoa e converte para JSON") 
		fmt.Println("2 - Chamada ViaCep")
		os.Exit(2)
		return
	}
	
	exemplo := os.Args[1]

	switch exemplo {
	case "1":
		Example1()
	case "2":
		Example2()
	default:
		fmt.Println("Exemplo inválido. Escolha 1 ou 2.")
		os.Exit(2)
		return
	}
}