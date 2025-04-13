# Trabalhando com JSON

Se você for trabalhar com desenvolvimento web com certeza você terá que lidar com APIs REST. E trabalhando com esta especificação, o JSON é o formato mais comum de enviar e receber dados.

Na linguagem Go, a forma mais simples e eficiente de trabalhar com JSON é através do pacote `enconding/json`, uma das bibliotecas padrão da linguagem. Este pacote fornece funções para converter dado JSON no formato de objetos do nosso programa, assim como transformar nossos objetos em JSON.

**observação**: Para converter um JSON em um objeto do programa, os atributos da struct precisam ser exportados, ou seja, seus nomes devem começar com uma letra maiúscula.

[Clique aqui para saber mais sobre JSON](https://github.com/cauelz/full-cycle-golang-expert/blob/master/chapter5/json/history.md)

# Principais funções

## `json.Marshal()`

Responsável por converter dados de diversos tipos para o formato JSON

`Sintaxe:`

```Go
json.Marshal(v interface{}) ([]byte, error)
```

`Parâmetros:`

- v: qualquer tipo de dado para converter em JSON (exemplo: structs, maps, slices entre outros)

`Retorno:`

- `[]byte`: retorna um array de bytes representando o JSON
- `error`: retorna um erro durante a conversão.

## `json.Unmarshal()`

Faz o caminho inverso: faz a leitura de um array de bytes que representa um JSON e converte para a estrutura indicada.

`Sintaxe:`

```Go
json.Unmarshal(data []byte, v interface{}) error
```

`Parâmetros:`

- data: um array de bytes que representa um JSON.
- v: qualquer tipo de dado para receber a conversão do JSON.

`Retorno:`

- `error`: retorna um erro durante a conversão.

# Trabalhando com TAGs, campos opcionais e customizações

Uma das funcionalidades mais interessantes da linguagem Go é o uso de `tags` em atributos de structs. Essas tags fornecem instruções adicionais para que as funções do pacote `json` trabalhem com mais flexibilidade.

Por exemplo, você pode usar tags para personalizar os nomes dos campos no JSON, ignorar campos ou torná-los opcionais. Veja um exemplo:

```Go
type Pessoa struct {
    Nome      string `json:"nome"`
    Idade     int    `json:"idade,omitempty"`
    Documento string `json:"-"`
}
```

- `json:"nome"`: define que o campo `Nome` será representado como `nome` no JSON.
- `json:"idade,omitempty"`: o campo `Idade` será omitido no JSON se estiver com o valor zero.
- `json:"-"`: o campo `Documento` será ignorado e não aparecerá no JSON.

Essas tags tornam o trabalho com JSON mais flexível e adaptável às necessidades de diferentes APIs e formatos de dados.

# Encoder e Decoder: trabalhando com fluxos de dados (streams)

Quando trabalhamos com dados muito grandes (como arquivos de gigabytes ou streams contínuos de dados), processar tudo de uma vez pode ser ineficiente e consumir muita memória. Para lidar com esses cenários, o pacote `encoding/json` fornece as interfaces `Encoder` e `Decoder`, que permitem trabalhar diretamente com fluxos de dados (`streams`), como arquivos, conexões de rede ou buffers.

[Clique aqui para saber mais sobre Streams](https://github.com/cauelz/full-cycle-golang-expert/blob/master/chapter5/streams/streams.md)

## `json.NewEncoder()`

O `Encoder` é usado para escrever objetos JSON diretamente em um fluxo de saída, como um arquivo, uma conexão de rede ou uma resposta HTTP. Ele é útil para gerar JSON de forma incremental, sem precisar carregar tudo na memória.

**Sintaxe:**

```Go
func NewEncoder(w io.Writer) *json.Encoder
```

**Exemplo de uso:**

```Go
package main

import (
    "encoding/json"
    "os"
)

type Pessoa struct {
    Nome      string `json:"nome"`
    Idade     int    `json:"idade,omitempty"`
    Documento string `json:"-"`
}

func main() {
    pessoa := Pessoa{
        Nome:      "João",
        Idade:     30,
        Documento: "123456789",
    }

    // Cria um arquivo para salvar o JSON
    file, err := os.Create("pessoa.json")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Cria um Encoder para escrever no arquivo
    encoder := json.NewEncoder(file)
    if err := encoder.Encode(pessoa); err != nil {
        panic(err)
    }

    // O JSON será salvo no arquivo "pessoa.json"
}
```

**Saída no arquivo `pessoa.json`:**

```json
{
    "nome": "João",
    "idade": 30
}
```

## `json.NewDecoder()`

O `Decoder` é usado para ler objetos JSON diretamente de um fluxo de entrada, como um arquivo, uma conexão de rede ou uma requisição HTTP. Ele é útil para processar JSON de forma incremental, especialmente quando o JSON é muito grande.

**Sintaxe:**

```Go
func NewDecoder(r io.Reader) *json.Decoder
```

**Exemplo de uso:**

```Go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Pessoa struct {
    Nome      string `json:"nome"`
    Idade     int    `json:"idade,omitempty"`
    Documento string `json:"-"`
}

func main() {
    // Abre o arquivo JSON
    file, err := os.Open("pessoa.json")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var pessoa Pessoa

    // Cria um Decoder para ler do arquivo
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&pessoa); err != nil {
        panic(err)
    }

    fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
}
```

**Saída no terminal:**

```
Nome: João, Idade: 30
```

## Funções adicionais do Encoder e Decoder

### `SetIndent` no Encoder

O método `SetIndent` permite formatar o JSON gerado com indentação, tornando-o mais legível.

**Exemplo:**

```Go
encoder := json.NewEncoder(os.Stdout)
encoder.SetIndent("", "  ") // Adiciona indentação de 2 espaços
encoder.Encode(pessoa)
```

**Saída:**

```json
{
  "nome": "João",
  "idade": 30
}
```

### `UseNumber` no Decoder

O método `UseNumber` permite que o `Decoder` trate números no JSON como o tipo `json.Number`, em vez de convertê-los automaticamente para `float64`. Isso é útil para preservar a precisão de números grandes.

**Exemplo:**

```Go
decoder := json.NewDecoder(file)
decoder.UseNumber()

var data map[string]interface{}
if err := decoder.Decode(&data); err != nil {
    panic(err)
}

fmt.Println(data["idade"].(json.Number).String()) // Preserva o número como string
```

### Decodificação incremental com `Token`

O método `Token` do `Decoder` permite processar JSON de forma incremental, token por token. Isso é útil para analisar grandes arquivos JSON sem carregá-los inteiramente na memória.

**Exemplo:**

```Go
file, _ := os.Open("pessoa.json")
defer file.Close()

decoder := json.NewDecoder(file)

for {
    token, err := decoder.Token()
    if err != nil {
        break
    }
    fmt.Printf("Token: %v\n", token)
}
```

**Saída:**

```
Token: {
Token: "nome"
Token: "João"
Token: "idade"
Token: 30
Token: }
```

## Vantagens do Encoder e Decoder

- **Eficiência de memória:** Trabalham diretamente com fluxos, evitando carregar todo o conteúdo na memória.
- **Flexibilidade:** Podem ser usados com qualquer tipo de `io.Reader` ou `io.Writer`, como arquivos, conexões de rede ou buffers.
- **Incremental:** Permitem processar JSON de forma incremental, ideal para grandes volumes de dados.

Essas ferramentas são indispensáveis para aplicações que lidam com grandes arquivos JSON ou streams contínuos de dados, como APIs ou sistemas de processamento em tempo real.
