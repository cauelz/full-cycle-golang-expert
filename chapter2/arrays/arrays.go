package arrays

import "fmt"

func Inicializar() {
	// Como definir um Array?

	var nomes [5]string

	nomes[0] = "Maria"
	nomes[1] = "João"
	nomes[2] = "José"
	nomes[3] = "Ana"
	nomes[4] = "Carlos"

	fmt.Print(nomes) // [Maria João José Ana Carlos]

	// Podemos definir um Array com valores iniciais

	var numeros = [5]int{1, 2, 3, 4, 5}
	fmt.Print(numeros) // [1 2 3 4 5]

	// Podemos definir um Array com valores iniciais e o compilador
	// irá inferir o tamanho do Array

	var frutas = [...]string{"Maçã", "Banana", "Pera"}
	fmt.Print(frutas) // [Maçã Banana Pera]

}
