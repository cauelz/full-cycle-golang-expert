package slices

import (
	"reflect"
	"testing"
)

func TestCriarSlices(t *testing.T) {
	resultado := CriarSlices()

	// Verifica se o slice retornado tem o tamanho esperado
	if len(resultado) != 5 {
		t.Errorf("CriarSlices(): esperado slice com tamanho 5, obtido %d", len(resultado))
	}

	// Verifica se os elementos são os esperados
	esperado := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("CriarSlices(): esperado %v, obtido %v", esperado, resultado)
	}
}

func TestManipularComAppend(t *testing.T) {
	// Como a função não retorna valores, vamos apenas verificar se ela executa sem erros
	ManipularComAppend()
}

func TestDemonstrarFatiamento(t *testing.T) {
	// Como a função não retorna valores, vamos testar a lógica de fatiamento separadamente
	slice := []int{1, 2, 3, 4, 5}

	// Testando fatiamento do meio
	meio := slice[1:3]
	if !reflect.DeepEqual(meio, []int{2, 3}) {
		t.Errorf("Fatiamento do meio: esperado [2 3], obtido %v", meio)
	}

	// Testando fatiamento do início
	inicio := slice[:2]
	if !reflect.DeepEqual(inicio, []int{1, 2}) {
		t.Errorf("Fatiamento do início: esperado [1 2], obtido %v", inicio)
	}

	// Testando fatiamento do fim
	fim := slice[3:]
	if !reflect.DeepEqual(fim, []int{4, 5}) {
		t.Errorf("Fatiamento do fim: esperado [4 5], obtido %v", fim)
	}
}

func TestCopiarSlices(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))
	copy(copia, original)

	if !reflect.DeepEqual(original, copia) {
		t.Errorf("CopiarSlices: esperado %v, obtido %v", original, copia)
	}

	// Verificar se é uma cópia real (modificar o original não afeta a cópia)
	original[0] = 999
	if copia[0] == 999 {
		t.Error("CopiarSlices: a modificação do slice original não deveria afetar a cópia")
	}
}

func TestRemoverElementos(t *testing.T) {
	// Testando a remoção de elementos
	slice := []int{1, 2, 3, 4, 5}
	esperado := []int{1, 2, 4, 5}

	slice = append(slice[:2], slice[3:]...)

	if !reflect.DeepEqual(slice, esperado) {
		t.Errorf("RemoverElementos: esperado %v, obtido %v", esperado, slice)
	}
}

func TestDemonstrarTamanhoCapacidade(t *testing.T) {
	slice := make([]int, 3, 5)

	if len(slice) != 3 {
		t.Errorf("Tamanho esperado: 3, obtido: %d", len(slice))
	}

	if cap(slice) != 5 {
		t.Errorf("Capacidade esperada: 5, obtida: %d", cap(slice))
	}
}

func TestIterarSobreSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	soma := 0
	somaEsperada := 15 // 1 + 2 + 3 + 4 + 5

	// Testando iteração com range
	for _, valor := range slice {
		soma += valor
	}

	if soma != somaEsperada {
		t.Errorf("Soma dos elementos: esperado %d, obtido %d", somaEsperada, soma)
	}

	// Testando iteração com for tradicional
	soma = 0
	for i := 0; i < len(slice); i++ {
		soma += slice[i]
	}

	if soma != somaEsperada {
		t.Errorf("Soma dos elementos (for tradicional): esperado %d, obtido %d", somaEsperada, soma)
	}
} 