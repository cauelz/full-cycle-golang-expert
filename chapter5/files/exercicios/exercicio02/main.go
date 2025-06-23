package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("arquivo.txt")

	if err != nil {
		fmt.Println("Erro ao criar arquivo: ", err)
	}

	_, err = io.WriteString(file, "Olá pessoal, este texto é para te desejar um otimo dia.")
	
	if err != nil {
		fmt.Println("Erro ao escrever arquivo: ", err)
		return
	}

	defer file.Close()
	
	fileLeitura, err := os.Open("arquivo.txt")

	if err != nil {
		fmt.Println("Erro ao abrir arquivo!")
		return
	}

	defer fileLeitura.Close()

	// content é um slice de bytes.
	content, err := io.ReadAll(fileLeitura)

	if err != nil {
		fmt.Println("erro ao ler arquivo: ", err)
		return
	}

	fmt.Println("Conteudo do arquivo: ", string(content))
}