# Pacotes em Go  

O conceito de pacotes (`packages`) em Go é bem direto e simples: trata-se de uma coleção de arquivos `.go` que estão no mesmo diretório. Uma boa prática é agrupar arquivos `.go` em um pacote que compartilhe a mesma finalidade ou contexto. Por exemplo: um pacote chamado "operações matemáticas" pode conter funções como soma, subtração, divisão etc.  

Todo arquivo que pertence a um pacote começa declarando o nome desse pacote com `package <nome>`.

---

## Pacote `main`  

O pacote `main` é especial, pois indica o **ponto de entrada** de um programa em Go. Quando compilamos e executamos um projeto, o Go procura pelo pacote `main` para iniciar a execução.  

Quando o pacote `main` é encontrado, é obrigatório ter uma função que marque o início da execução do programa: a `func main()`. Essa função não recebe argumentos e não retorna valores.  

Exemplo:
```go
package main  

func main() {  
    // Código de inicialização aqui  
}
```
## Estrutura dos Pacotes

Cada diretório em Go é considerado um pacote independente. No exemplo abaixo, temos um projeto chamado `meu-projeto`, onde também definimos um módulo neste projeto para o gerenciamento das dependências. Abordaremos mais detalhes sobre módulos nos próximos capítulos.

```text
meu-projeto/
├── go.mod          # Módulo (gerenciamento de dependências)
└── pkg/
    └── saudacao/   # Pacote saudacao
        ├── saudacao.go
        └── saudacao_test.go
```

O nome do pacote é definido dentro de cada arquivo `.go` e, como boa prática, costumamos utilizar o mesmo nome do diretório.

Para importar um pacote em outro arquivo pertencente a um pacote diferente, utiliza-se o nome do módulo definido no arquivo `go.mod` seguido do caminho para o pacote. No exemplo, se o nome do módulo for `meu-projeto`, a importação ficaria assim:

```go

import "meu-projeto/pkg/saudacao"

```

## Visibilidade de Identificadores

Em Go, para conseguirmos tornar elementos e estruturas do nosso pacote acessíveis externamente, a linguagem utiliza a capitalização da primeira letra de cada Estrutura (Arrays, Variaveis, Slices, Maps, Structs e Interfaces). Se a letra for maiuscula, esta estrutura é considerada pública; caso contrário, a estrutura é privada.

Exemplo:

```go

package saudacao

var saudacao string = "Olá Gopher!"

func Saudar() {
    println(saudacao)
}

```
Neste exemplo, a função `Saudar()` é publica. Enquanto a variavel `saudacao` é privada.

## Uso de Pacotes

No capítulo 2, você percebeu que utilizamos bastante o pacote `fmt`, que possui funções para o output de informações. Vamos lembrar como importamos o pacote?

Usamos a palavra reservada `import` mais o nome do pacote.

```Go

package saudacao

import "fmt"

func Saudacao() {
    fmt.Println("Olá Go!")
}

```

### Alias para pacotes importados

Podemos definir um alias para utilizarmos os pacotes.

```Go
import (
    s "meu-projeto/pkg/saudacao"
)
```

### Importação Anônima

```Go

import _ "github.com/lib/pq" // neste caso, só queremos inicializar o driver do postgreSQL

```

## Escopo e Inicialização de pacotes

Podemos declarar uma função chamada `init()` dentro dos pacotes. Elas são executadas automaticamente antes da função main(), na ordem de importação.

```Go
package saudacao

import "fmt"

func init() {
    fmt.Println("Pacote inicializado!")
}
```

# Módulos em Go

