package loops

import "fmt"

func Inicializar() {
	// Como definir um loop?

	frutas := [...]string{"Maçã", "Banana", "Pera"}

	for i := 0; i < len(frutas); i++ {
		fmt.Println(frutas[i])
	}

	// Podemos usar o range para iterar sobre um Array

	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}

	// Podemos usar o range para iterar sobre um Array e ignorar o índice

	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// Podemos o range para iterar sobre um Slices
	slice := []string{"Maçã", "Banana", "Pera"}

	for i, fruta := range slice {
		fmt.Println(i, fruta)
	}

	// Podemos usar o range para iterar sobre um Map

	mapa := map[string]string{
		"nome":    "Maria",
		"idade":   "30",
		"cidade":  "São Paulo",
		"estado":  "SP",
	}

	for chave, valor := range mapa {
		fmt.Println(chave, valor)
	}

	// Podemos usar o range para iterar sobre um String

	str := "Olá Mundo"

	for i, letra := range str {
		fmt.Println(i, letra)
	}
	
}