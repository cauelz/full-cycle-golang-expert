
# Primeiro Capítulo - Variaveis e Tipos

Neste primeiro capítulo vamos compreender como declarar variaveis em GO assim como seus principais tipos. GO é uma linguagem fortemente tipada, ou seja, todas as variaveis que declaramos no código precisam indicar qual tipo de dado que elas vão armazenas como Texto, Inteiro, Boleano entre outros.  

## Usando a palavra chave "var" (declaração explícita):

Exemplo 1:

````
// Aqui temos três variaveis utilizando a palavra chave "var":

// nome do tipo texto
// idade do tipo inteiro
// ehMaiorDeIdade do tipo boleano
// preco do tipo ponto flutuante

//var nome string
//var idade int
//var ehMaiorDeIdade bool
//var preco float64 = 30.00

func variaveis() {
    println(nome)
    println(idade)
    println(ehMaiorDeIdade)
}

````

Veja que a variável "preco" foi declarada e inicializada, ou seja, ela "nasceu" com um valor pré definidos pos nós.

DESAFIO: Tente supor qual os valores que as variaveis nome, idade, ehMaiorDeIdade e preco vão imprimir no console.

## Inferência de tipo com ":=" (short hand):

`````
// Podemos declarar variaveis de maneira mais curta e rápida

func shortHand() {
    nome := "John""
    sobrenome := "Doe""
    idade := 30
}

`````
Desta forma, o compilador "deduz" o tipo da variavel de acordo com o valor atribuido a ela. Esta forma só pode ser utilizada em funções.

## Declaração de multiplas variáveis:

`````
// Podemos declarar multiplas variáveis utilizando "var" e ":=""

var (
    nome string = "Carlos""
    idade int = 30
    preco float64 = 30.00
)

func shortHand() {
    nome, idade, preco := "Carlos", 30, 30.00
}

`````


