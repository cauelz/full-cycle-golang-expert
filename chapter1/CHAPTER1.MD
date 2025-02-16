
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

Veja que a variavel "preco" foi declarada e inicializada, ou seja, ela "nasceu" com um valor pré definidos pos nós.

DESAFIO: Tente supor qual os valores que as variaveis nome, idade, ehMaiorDeIdade e preco vão imprimir no console.