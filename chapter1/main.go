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

func main() {
	valoresPadrão()
}

func valoresPadrão() {
	println("Valores padrão em GO")
	println(nome)
	println(idade)
	println(ehMaiorDeIdade)
	println(preco)
}