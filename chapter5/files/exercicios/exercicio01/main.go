package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("mensagem.txt")

	if err != nil {
		fmt.Println("Erro ao criar arquivo: ", err)
		return
	}

	defer file.Close()

	_, err = io.WriteString(file, "Olá, este é meu primeiro arquivo em Go!")

	if err != nil {
		fmt.Println("Erro ao escrever no arquivo: ", err)
		return
	}

	fmt.Println("Arquivo escrito com sucesso.")

}