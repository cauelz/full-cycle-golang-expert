package slices

import (
	"fmt"
)

// CriarSlices demonstra diferentes formas de criar slices
func CriarSlices() []int {
	// Declarando um slice vazio
	var slice1 []int
	_ = slice1 // Explicitamente ignorando para fins didáticos

	// Declarando e inicializando um slice
	slice2 := []int{1, 2, 3, 4, 5}

	// Criando slice com make (tipo, tamanho, capacidade)
	slice3 := make([]int, 3, 5)
	_ = slice3 // Explicitamente ignorando para fins didáticos

	return slice2 // retornando para uso em outros exemplos
}

// ManipularComAppend demonstra como adicionar elementos em slices
func ManipularComAppend() {
	slice1 := []int{}
	slice2 := []int{1, 2, 3}

	// Adicionando elementos com append
	slice1 = append(slice1, 10)
	slice2 = append(slice2, 4, 5)
}

// DemonstrarFatiamento mostra como fatiar slices
func DemonstrarFatiamento() {
	original := []int{1, 2, 3, 4, 5}
	
	// Pega elementos do índice 1 até o 3 (não inclui o 3)
	meio := original[1:3]     // [2, 3]
	inicio := original[:2]    // [1, 2]
	fim := original[3:]       // [4, 5]

	_ = meio
	_ = inicio
	_ = fim
}

// CopiarSlices demonstra como copiar slices
func CopiarSlices() {


	original := []int{1, 2, 3, 4, 5}
	
	// Para copiar um slice, podemos usar a função copy.
	// A função copy retorna o número de elementos copiados e não retorna um novo slice.
	copia := make([]int, len(original))
	copy(copia, original)
	
}

// RemoverElementos demonstra como remover elementos de um slice
func RemoverElementos() {
	slice := []int{1, 2, 3, 4, 5}
	
	// Remove o elemento do índice 2
	slice = append(slice[:2], slice[3:]...)
}

// DemonstrarTamanhoCapacidade mostra como verificar tamanho e capacidade
func DemonstrarTamanhoCapacidade() {
	slice := make([]int, 3, 5)
	
	tamanho := len(slice)
	capacidade := cap(slice)

	fmt.Printf("Tamanho: %d, Capacidade: %d\n", tamanho, capacidade)
}

// IterarSobreSlice demonstra diferentes formas de iterar sobre um slice
func IterarSobreSlice() {
	slice := []int{1, 2, 3, 4, 5}

	// Usando range
	for i, valor := range slice {
		fmt.Printf("Índice: %d, Valor: %d\n", i, valor)
	}

	// Usando for tradicional
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Valor na posição %d: %d\n", i, slice[i])
	}
}

// Inicializar chama todas as funções de demonstração
func Inicializar() {
	CriarSlices()
	ManipularComAppend()
	DemonstrarFatiamento()
	CopiarSlices()
	RemoverElementos()
	DemonstrarTamanhoCapacidade()
	IterarSobreSlice()
}