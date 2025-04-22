# Trabalhando com Templates em Go a partir de Arquivos

Este exemplo demonstra como trabalhar com templates HTML em Go utilizando múltiplos arquivos de template, incluindo conceitos avançados como template composition e funções customizadas.

## Estrutura do Projeto

```
09-templates-with-file/
├── main.go         # Arquivo principal com a lógica do servidor
├── header.html     # Template do cabeçalho HTML
├── content.html    # Template principal com o conteúdo
├── footer.html     # Template do rodapé HTML
└── go.mod         # Arquivo de módulo Go
```

## Conceitos Demonstrados

1. Composição de templates (template composition)
2. Funções customizadas em templates
3. Manipulação de dados estruturados
4. Loops em templates
5. Tratamento de erros

## Estrutura dos Arquivos

### 1. Estruturas de Dados (main.go)

```go
type Curso struct {
    Nome         string
    CargaHoraria int
}

type Cursos []Curso
```

Esta estrutura define o formato dos dados que serão exibidos no template:
- `Nome`: Nome do curso
- `CargaHoraria`: Duração do curso em horas
- `Cursos`: Slice de `Curso` para armazenar múltiplos cursos

### 2. Templates HTML

#### header.html
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Veja nossos cursos disponíveis</title>
</head>
<body>
```

#### content.html
```html
{{template "header.html"}}

<h1>Cursos</h1>

<table>
    <thead>
        <tr>
            <th>Nome</th>
            <th>Horas</th>
        </tr>
    </thead>
    <tbody>
        {{ range . }}
        <tr>
            <td>{{ .Nome | ToUpper}}</td>
            <td>{{ .CargaHoraria}}</td>
        </tr>
        {{ end }}
    </tbody>
</table>

{{template "footer.html"}}
```

#### footer.html
```html
</body>
</html>
```

### 3. Implementação do Servidor (main.go)

```go
package main

import (
    "html/template"
    "net/http"
    "strings"
)

func main() {
    // Define a lista de templates que serão utilizados
    var templates []string = []string{
        "content.html",
        "header.html",
        "footer.html",
    }
    
    // Configura o handler para a rota principal
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        // Cria um novo template baseado no arquivo principal
        t := template.New("content.html")
        
        // Adiciona função customizada para converter texto para maiúsculas
        t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})

        // Carrega e parseia todos os arquivos de template
        tmp := template.Must(t.ParseFiles(templates...))
        
        // Executa o template com os dados dos cursos
        error := tmp.Execute(w, Cursos{
            {"Go", 40},
            {"Java", 90},
            {"Javascript", 60},
        })
    
        // Tratamento de erro
        if error != nil {
            panic(error)
        }
    })

    // Inicia o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}
```

## Explicação Detalhada

### 1. Composição de Templates

O exemplo utiliza três arquivos de template separados:
- `header.html`: Contém a estrutura inicial do HTML
- `content.html`: Contém o conteúdo principal e referencia os outros templates
- `footer.html`: Contém o fechamento do HTML

Esta separação permite:
- Melhor organização do código
- Reutilização de componentes
- Manutenção mais fácil

### 2. Funções Customizadas

O exemplo demonstra como adicionar funções customizadas aos templates:
```go
t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
```

No template, a função é usada com o operador pipe:
```html
{{ .Nome | ToUpper}}
```

### 3. Carregamento de Templates

```go
var templates []string = []string{
    "content.html",
    "header.html",
    "footer.html",
}
```

Todos os arquivos de template são carregados usando `ParseFiles`:
```go
tmp := template.Must(t.ParseFiles(templates...))
```

### 4. Sintaxe do Template

O arquivo `content.html` demonstra várias funcionalidades:

1. **Inclusão de outros templates**:
   ```html
   {{template "header.html"}}
   ```

2. **Iteração sobre dados**:
   ```html
   {{ range . }}
   <tr>
       <td>{{ .Nome | ToUpper}}</td>
       <td>{{ .CargaHoraria}}</td>
   </tr>
   {{ end }}
   ```

3. **Uso de funções com pipe**:
   ```html
   {{ .Nome | ToUpper}}
   ```

### 5. Tratamento de Erros

O código utiliza `template.Must` para tratamento de erros durante o parsing:
```go
tmp := template.Must(t.ParseFiles(templates...))
```

E também verifica erros durante a execução:
```go
error := tmp.Execute(w, Cursos{...})
if error != nil {
    panic(error)
}
```

## Como Executar

1. Certifique-se de que todos os arquivos estão no mesmo diretório
2. Execute o servidor:
   ```bash
   go run main.go
   ```
3. Acesse `http://localhost:8080` no navegador

## Resultado

Ao acessar o servidor, você verá uma tabela HTML formatada com:
- Nomes dos cursos em maiúsculas (devido à função ToUpper)
- Carga horária de cada curso
- Layout HTML completo com header e footer

## Dicas e Boas Práticas

1. **Organização de Templates**:
   - Mantenha templates relacionados em arquivos separados
   - Use nomes descritivos para os arquivos
   - Considere uma estrutura de diretórios para projetos maiores

2. **Funções Customizadas**:
   - Crie funções para lógica de apresentação complexa
   - Mantenha as funções simples e focadas
   - Documente o comportamento esperado

3. **Tratamento de Erros**:
   - Sempre verifique erros durante o parsing
   - Implemente tratamento adequado de erros em produção
   - Use logging apropriado

4. **Performance**:
   - Em produção, considere carregar templates apenas uma vez na inicialização
   - Implemente cache quando apropriado
   - Monitore o uso de memória com grandes conjuntos de dados 