package main

//var nome string
//var idade int
//var ehMaiorDeIdade bool
//var preco float64

var (
	nome string
	idade int
	ehMaiorDeIdade bool
	preco float64
)

const constante = "Valor constante"

type Idade int

func main() {
	valoresPadrao()
	shortHand()
	constants()
	newAge()
}

func valoresPadrao() {
	println("Valores padrão em GO")
	println(nome)
	println(idade)
	println(ehMaiorDeIdade)
	println(preco)
}

func shortHand() {
	nome, idade, ativo := "Carlos", 30, true

	println("Short Hand")
	println(nome)
	println(idade)
	println(ativo)
}

func constants() {
	println("Constantes", constante)
}

func newAge() {
	var novaIdade Idade = 30
	println("Nova idade", novaIdade)
}