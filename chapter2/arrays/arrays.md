# Arrays em Golang

Arrays são estruturas de dados fundamentais e muito comuns em várias linguagens de programação. Em Go, um array é uma coleção de elementos do mesmo tipo, armazenados de forma sequencial e acessados por um índice numérico, que sempre começa em zero.

A principal característica dos arrays em Golang é que eles possuem tamanho fixo, definido no momento da declaração. Ou seja, após definir um array com tamanho 3, ele sempre terá exatamente 3 posições, nem mais, nem menos.

Por exemplo, imagine que queremos armazenar os nomes: Lucas, Gabriela e João. Podemos fazer isso da seguinte forma:

````
// Exemplo de declaração e atribuição de valores do tipo texto em um Array

var nomes [3]string

nomes[0] = "Lucas"
nomes[1] = "Gabriela""
nomes[2] = "João"

fmt.Println("nome: ", nome)

// Exemplo de declaração e inicialização

idades := [3]int{10, 33, 5}

fmt.Println("idades: ", idades)

````

Nesse exemplo acima:

- Declaramos um array chamado nomes com um tamanho de 3 posições.
- Atribuímos um nome para cada posição do array.
- Imprimimos todos os nomes armazenados no array com o pacote "fmt"

## Utilização de reticências para o compilador inferir o tamanho de um Array

Podemos utilizar reticências para dizer para o compilador "indicar" o tamanho do array automaticamente com base na quantidade de elementos fornecidos na inicialização.

Muito útil quando precisamos adicionar mais elementos em um array sem precisar ficar "contando" elementos.

observação: mesmo que o tamanho seja definido "magicamente", o array ainda possui seu tamanho fixo.

`````

// Exemplo do uso de reticências

nomes := [...]string{"João", "Paulo", "Gabriela"}


`````
