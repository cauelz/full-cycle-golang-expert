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

