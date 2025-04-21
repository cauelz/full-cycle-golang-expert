# Templates em Go

Os Templates em Go são uma ferramenta poderosa para gerar saídas textuais baseadas em modelos predefinidos. Eles são especialmente úteis para gerar HTML, documentos e outros formatos de texto onde você precisa combinar dados dinâmicos com estruturas fixas.

## Pacotes de Templates

Go oferece dois pacotes principais para templates:

1. `text/template`: Para templates de texto genéricos
2. `html/template`: Específico para HTML, com proteções contra vulnerabilidades como XSS

## Características Principais

- **Sintaxe Simples**: Usa {{}} para delimitar ações e expressões
- **Segurança**: Escape automático de caracteres especiais (especialmente no html/template)
- **Composição**: Permite reutilização através de templates aninhados
- **Funções**: Suporta funções built-in e customizadas
- **Controle de Fluxo**: Oferece if/else, range, with e outros controles

## Exemplo Básico

```go
package main

import (
    "html/template"
    "os"
)

type Produto struct {
    Nome  string
    Preco float64
}

func main() {
    // Define o template
    tmpl := `
    <h1>Produto: {{.Nome}}</h1>
    <p>Preço: R$ {{.Preco}}</p>
    `

    // Cria um novo template
    
    t := template.Must(template.New("produto").Parse(tmpl))

    // Dados para o template
    produto := Produto{
        Nome:  "Notebook",
        Preco: 2999.99,
    }

    // Executa o template
    err := t.Execute(os.Stdout, produto)
    if err != nil {
        panic(err)
    }
}
```

Este exemplo demonstra:
1. Como criar uma estrutura de dados
2. Como definir um template com marcadores {{}}
3. Como parsear o template
4. Como executar o template com dados

## Função template.Must

A função `template.Must` é um wrapper que serve para dois propósitos principais:

1. **Tratamento de Erros**: Ela recebe um template e um erro como argumentos (`template.Must(template, error)`). Se o erro não for nil, ela causa um panic imediatamente.
2. **Inicialização Segura**: É especialmente útil durante a inicialização do programa, garantindo que os templates sejam válidos antes da execução continuar.

Exemplo de uso:
```go
// Sem Must - Tratamento manual de erro
t, err := template.New("exemplo").Parse("Hello {{.Name}}")
if err != nil {
    // tratar erro
}

// Com Must - Panic automático em caso de erro
t := template.Must(template.New("exemplo").Parse("Hello {{.Name}}"))
```

Use `template.Must` quando:
- Estiver definindo templates durante a inicialização do programa
- Tiver certeza que o template é válido e quiser um código mais conciso
- Quiser que o programa falhe rapidamente se houver um erro no template

## Exemplo Completo com template.Must

```go
package main

import (
    "fmt"
    "html/template"
    "os"
)

// Estrutura de dados para o template
type Usuario struct {
    Nome  string
    Email string
    Idade int
}

func main() {
    // Exemplo 1: Template válido com Must
    templateValido := `
        <div>
            <h1>Usuário: {{.Nome}}</h1>
            <p>Email: {{.Email}}</p>
            {{if ge .Idade 18}}
                <p>Maior de idade</p>
            {{else}}
                <p>Menor de idade</p>
            {{end}}
        </div>
    `
    
    // Usando Must com template válido
    tmpl := template.Must(template.New("usuario").Parse(templateValido))
    
    usuario := Usuario{
        Nome:  "João",
        Email: "joao@email.com",
        Idade: 25,
    }
    
    // Executa o template
    fmt.Println("Executando template válido:")
    tmpl.Execute(os.Stdout, usuario)
    
    // Exemplo 2: Template inválido com Must (causará panic)
    templateInvalido := `
        <div>
            <h1>Usuário: {{.Nome}</h1> {{/* Falta fechar a chave */}}
        </div>
    `
    
    // Esta linha causará panic devido ao erro de sintaxe
    // tmpl = template.Must(template.New("invalido").Parse(templateInvalido))
    
    // Exemplo 3: Forma segura sem Must
    if t, err := template.New("seguro").Parse(templateInvalido); err != nil {
        fmt.Printf("Erro ao parsear template: %v\n", err)
    }
}
```

Este exemplo demonstra:
1. Como usar `template.Must` com um template válido
2. O que acontece com um template inválido (comentado para evitar o panic)
3. Como fazer o tratamento manual de erro sem usar Must
4. Uso de condicionais dentro do template com `{{if}}` e funções como `ge` (greater or equal)





