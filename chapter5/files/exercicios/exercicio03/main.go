package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("mensagem.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}

	defer file.Close()

	message := "Esta é uma nova linha adicionada ao arquivo.\n"


	_, err = file.Write([]byte(message))

	if err != nil {
		fmt.Println("Erro ao escrever messagem no arquivo: ", err)
		return
	}

}
