package main

import (
	"bufio" // Pacote para leitura eficiente de arquivos linha a linha
	"fmt"
	"os"
)

func main() {
	// Abre o arquivo mensagem.txt para leitura
	file, err := os.Open("mensagem.txt")
	if err != nil {
		fmt.Println("Erro ao abrir arquivo: ", err)
		return
	}
	// Garante que o arquivo será fechado ao final da função
	defer file.Close()

	// Cria um scanner para ler o arquivo linha a linha
	// É possível passar o ponteiro 'file' porque *os.File implementa a interface io.Reader,
	// que é o que o bufio.NewScanner espera como argumento.
	scanner := bufio.NewScanner(file)

	counter := 0 // Variável para contar o número de linhas

	// O método Scan avança o scanner para a próxima linha
	for scanner.Scan() {
		counter++ // Incrementa o contador a cada linha lida
	}

	// Verifica se houve algum erro durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura do arquivo:", err)
		return
	}

	// Exibe o número total de linhas
	fmt.Println(counter)
}
