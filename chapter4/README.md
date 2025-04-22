# Pacotes e Módulos em Go

Este capítulo aborda os conceitos fundamentais de pacotes e módulos em Go, incluindo exemplos práticos e boas práticas de organização de código.

## Índice
1. [Conceitos Básicos de Pacotes](#conceitos-básicos-de-pacotes)
2. [Estrutura e Organização](#estrutura-e-organização)
3. [Visibilidade e Escopo](#visibilidade-e-escopo)
4. [Importação e Uso](#importação-e-uso)
5. [Módulos](#módulos)
6. [Boas Práticas](#boas-práticas)

## Conceitos Básicos de Pacotes

### O que é um Pacote?
Um pacote em Go é uma coleção de arquivos `.go` localizados no mesmo diretório que compartilham a mesma declaração `package`. Os pacotes são a unidade fundamental de organização de código em Go, permitindo:
- Reutilização de código
- Encapsulamento
- Modularização
- Gerenciamento de dependências

### Pacote Main
O pacote `main` tem um papel especial em Go:
- É o ponto de entrada de aplicações executáveis
- Deve conter uma função `main()`
- Não pode ser importado por outros pacotes

```go
package main

func main() {
    // Ponto de entrada do programa
}
```

## Estrutura e Organização

### Estrutura Básica de um Projeto
```
meu-projeto/
├── go.mod          # Definição do módulo e dependências
├── main.go         # Arquivo principal (package main)
├── internal/       # Código privado do projeto
│   └── config/
│       └── config.go
└── pkg/           # Código público reutilizável
    ├── database/
    │   └── db.go
    └── utils/
        └── helper.go
```

### Convenções de Nomenclatura
1. **Nomes de Pacotes**:
   - Usar nomes curtos e descritivos
   - Evitar underscores ou camelCase
   - Usar substantivos, não verbos
   ```go
   package database  // Bom
   package db_utils  // Evitar
   ```

2. **Nomes de Arquivos**:
   - Usar snake_case para nomes compostos
   - Sufixo `_test.go` para arquivos de teste
   ```
   user_repository.go
   user_repository_test.go
   ```

## Visibilidade e Escopo

### Regras de Visibilidade
1. **Exportação (Público)**:
   - Identificadores que começam com letra maiúscula
   - Acessíveis fora do pacote
   ```go
   type User struct {     // Público
       Name string        // Público
       email string      // Privado
   }
   ```

2. **Não Exportação (Privado)**:
   - Identificadores que começam com letra minúscula
   - Acessíveis apenas dentro do pacote

### Escopo de Pacote
```go
package math

var pi = 3.14159  // variável de pacote (privada)
const Pi = 3.14159 // constante exportada (pública)

func Calculate() float64 {  // função exportada
    return internalCalc()   // função privada
}

func internalCalc() float64 {
    return pi * 2
}
```

## Importação e Uso

### Formas de Importação
1. **Importação Simples**:
   ```go
   import "fmt"
   ```

2. **Importação Múltipla**:
   ```go
   import (
       "fmt"
       "strings"
       "time"
   )
   ```

3. **Importação com Alias**:
   ```go
   import (
       f "fmt"
       s "strings"
   )
   ```

4. **Importação com Ponto**:
   ```go
   import . "math"  // Não recomendado exceto em testes
   ```

5. **Importação Blank**:
   ```go
   import _ "github.com/lib/pq"  // Apenas para efeitos colaterais
   ```

### Inicialização de Pacotes

1. **Função init()**:
   - Executada antes da main()
   - Pode haver múltiplas por pacote
   - Ordem de execução determinística

```go
package database

func init() {
    // Inicialização do pacote
}

func init() {
    // Outra inicialização
}
```

## Módulos

### O que são Módulos?
Módulos são a unidade de distribuição de código em Go, introduzidos na versão 1.11. Eles permitem:
- Versionamento de dependências
- Reprodutibilidade de builds
- Gerenciamento de dependências

### Criando um Módulo
```bash
go mod init meu-projeto
```

### Estrutura do go.mod
```
module meu-projeto

go 1.20

require (
    github.com/pkg/errors v0.9.1
    golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
)
```

## Boas Práticas

1. **Organização de Código**:
   - Um pacote por diretório
   - Nomes de pacotes = nomes de diretórios
   - Separar código por responsabilidade

2. **Interfaces**:
   - Definir interfaces onde são usadas
   - Manter interfaces pequenas
   ```go
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   ```

3. **Documentação**:
   - Documentar pacotes públicos
   - Usar comentários no estilo godoc
   ```go
   // Package database provides database connection utilities.
   package database
   
   // Connect establishes a new database connection.
   func Connect() error {
       // ...
   }
   ```

4. **Testes**:
   - Arquivos de teste no mesmo diretório
   - Nomear arquivos com sufixo _test.go
   ```go
   // math_test.go
   package math_test  // teste externo
   
   func TestAdd(t *testing.T) {
       // ...
   }
   ```

## Exemplos Práticos

### 1. Criando um Pacote Utilitário
```go
// pkg/utils/string.go
package utils

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### 2. Usando o Pacote
```go
// main.go
package main

import (
    "fmt"
    "meu-projeto/pkg/utils"
)

func main() {
    reversed := utils.Reverse("Hello, Go!")
    fmt.Println(reversed)
}
```

## Recursos Adicionais

1. **Documentação Oficial**:
   - [Go Packages](https://golang.org/pkg/)
   - [Go Modules](https://blog.golang.org/using-go-modules)

2. **Ferramentas**:
   - `go mod tidy`: Atualiza dependências
   - `go doc`: Documentação local
   - `godoc`: Servidor de documentação

3. **Comandos Úteis**:
   ```bash
   go mod init     # Inicializa um novo módulo
   go mod tidy     # Atualiza dependências
   go mod vendor   # Cria diretório vendor
   go list -m all  # Lista todas as dependências
   ```

# Pacotes em Go  

O conceito de pacotes (`packages`) em Go é bem direto e simples: trata-se de uma coleção de arquivos `.go` que estão no mesmo diretório. Uma boa prática é agrupar arquivos `.go` em um pacote que compartilhe a mesma finalidade ou contexto. Por exemplo: um pacote chamado "operações matemáticas" pode conter funções como soma, subtração, divisão etc.  

Todo arquivo que pertence a um pacote começa declarando o nome desse pacote com `package <nome>`.

---

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