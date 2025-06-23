package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Abre o arquivo de origem para leitura
	file, err := os.Open("mensagem.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}
	defer file.Close() // Garante o fechamento do arquivo de origem

	// Cria o arquivo de destino para escrita
	newFile, err := os.Create("copia.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo: ", err)
		return
	}
	defer newFile.Close() // Garante o fechamento do arquivo de destino

	// Lê todo o conteúdo do arquivo de origem
	byteSlice, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Erro ao ler arquivo: ", err)
		return
	}

	// Escreve o conteúdo lido no novo arquivo
	_, err = newFile.Write(byteSlice)
	if err != nil {
		fmt.Println("Erro ao escrever arquivo: ", err)
		return
	}

	fmt.Println("Arquivo copiado com sucesso!")
}
