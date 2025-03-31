# Primeiro Capítulo - Variaveis e Tipos

Neste primeiro capítulo vamos compreender como declarar variaveis e constantes em GO assim como seus principais tipos. GO é uma linguagem fortemente tipada, ou seja, todas as variaveis declaradas no código precisam indicar qual tipo de dado que elas vão armazenar. Os tipos mais comuns são: texto, inteiro, boleano e ponto-flutuante.  

## Usando a palavra chave "var" (declaração explícita):

Exemplo 1:

````
// Aqui temos três variaveis utilizando a palavra chave "var":

// nome do tipo texto
// idade do tipo inteiro
// ehMaiorDeIdade do tipo boleano
// preco do tipo ponto-flutuante

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

Exemplo 2:

`````
// Podemos declarar variaveis de maneira mais curta e rápida

func shortHand() {
    nome := "John""
    sobrenome := "Doe""
    idade := 30
}

`````
Desta forma, o compilador "deduz" o tipo da variavel de acordo com o valor atribuido a ela. Esta forma só pode ser utilizada em funções, excluindo o escopo global.

## Declaração de multiplas variáveis:

Exemplo 3:

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

## Declaração de Constantes:

Exemplo 4:

`````
// Podemos declarar constantes utilizando a palavra "const""

const IDADE_LIMITE = 18

func constant() {
    println("Idade Limite", IDADE_LIMITE)
}

`````

Tente alterar o valor de IDADE_LIMITE dentro da função constant e veja o que acontece.

## Criação e declaração de novos tipos:

`````
// A linguagem Go nos permite definir nossos próprios tipos utilizando a palavra-chave type.


type Idade int

func createAge() {
    var idade Idade = 10

    println("idade", idade)
}

`````
Isso possibilita deixar o codigo mais expressivo e fácil de manter.

## Saída de Dados com o pacote "fmt"

O Pacote "fmt" nos permite imprimir valores no terminal, facilitando a visualização e a depuração das informações do nosso código. É um dos pacotes mais importantes e simples da linguagem.

Vamos explorá-lo com mais detalhes no futuro.

### Funcionalidades Principais

1. fmt.Print
    - imprime valores sem quebra de linha

2. fmt.Println
    - imprime valores com quebra de linha

3. fmt.Printf
    - imprime valores formatados a partir de variaveis.

`````
fmt.Print("Texto sem quebra de linha")
fmt.Println("Texto com quebra de linha")
fmt.Printf("A pessoa chamada %s tem %d anos de idade!", "John", 30)
`````
