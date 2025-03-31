# Loops em Golang

A simplicidade é uma das marcas registradas da linguagem Go, como mencionado na introdução deste repositório. Isso se torna evidente quando trabalhamos com iterações, diferentemente de outras linguagens, em Go há apenas uma estrutura para percorrer coleções – o clássico for.

Para desenvolvedores acostumados a linguagens com paradigmas funcionais, a ausência de métodos embutidos como map, filter ou reduce pode parecer uma limitação. Esses recursos, que em outras linguagens trazem lógica pré-definida para iterações, são intencionalmente omitidos em Go. A filosofia da linguagem prioriza clareza e controle explícito – você escreve exatamente o que o código faz, sem abstrações ocultas.

Mas não se engane! O for em Go é surpreendentemente versátil. Seja em sua forma tradicional, como range para iterar sobre slices/mapas, ou em implementações customizadas (como while simulados), essa única palavra-chave oferece toda a flexibilidade necessária. A simplicidade de Go está justamente em reduzir opções para aumentar a legibilidade e a manutenibilidade do código.

## Resumo das Formas de Usar `for` em Go

| Formato                   | Uso Típico                      |
|---------------------------|----------------------------------|
| `for i := 0; i < N; i++`  | Iteração tradicional             |
| `for condição { ... }`    | Simula `while`                   |
| `for { ... }`             | Loop infinito                    |
| `for range coleção { ... }` | Iteração sobre coleções         |

## Exemplo 1: Forma básica de um For Loop em Go.

`````
for i := 0; i < 5; i++ {
    fmt.Println("Valor de i: ", i)
} 

`````

## Exemplo 2: For como While

`````
i := 0

for i < 10 {
    fmt.Println("Valor de i: ", i)
    i++
}

`````

## Exemplo 3: Loop infinito

`````
for {
    fmt.Prinln("Um dia eu termino...")
}
`````

## Exemplo 4: Interando sobre Coleções

Em Go, percorrer coleções pode ser feita atravez da palavra reservada "range". Ela identifica automaticamente a estrutura de dados e percorre todos seus elementos. Podemos utilizar esta forma com arrays, slices, maps(ordem de iteração não é garantida) e strings.

O "range" retorna dois valores em cada iteração: índice/chave e valor.

Caso um dos valores não seja necessário, você pode ignorá-los usando "_" (blank identifier).

`````

// Arrays e Slices

frutas := [4]string{"banana", "laranja", "abacaxi", "maçã"}

for indice, valor := range frutas {
    fmt.Printf("Posição %d: %s \n")
}

// apenas valores ignorando o índice

for _, valor := range frutas {
    fmt.Println("Frutas: ", valor)
}

// Maps

cores := map[string]string{  
    "vermelho": "#ff0000",  
    "verde": "#00ff00",  
}  

// Iteração por chave e valor:  
for chave, hex := range cores {  
    fmt.Printf("Cor: %s → Código: %s\n", chave, hex)  
} 

// Strings

palavra := "Go"  

// Itera sobre os caracteres (runas):  
for i, runa := range palavra {  
    fmt.Printf("Posição %d: Unicode %U → Caractere '%c'\n", i, runa, runa)  
}  

`````

## Exemplo 5: Controle de fluxo com "break" e "continue"

`````
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Pula para a próxima iteração
    }
    if i == 7 {
        break // Sai do loop
    }
    fmt.Println("Número ímpar:", i)
}

`````
