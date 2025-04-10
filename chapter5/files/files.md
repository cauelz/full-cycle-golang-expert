# Manipulação de Arquivos e Diretórios em Go.

Sempre tive muita dificuldade para entender como funcionam as operações com arquivos em linguagens de programação. Por isso, neste capítulo, vou me esforçar para que cada explicação seja didática, ajudando você a dominar esses conceitos de forma simples.

Podemos resumir a essência deste capítulo em dois verbos fundamentais: ler e escrever. Essas duas ações são a base de qualquer interação com arquivos e diretórios. Para realizá-las em Go, utilizamos principalmente dois pacotes da biblioteca padrão:

`os`: Responsável por operações de baixo nível com o sistema operacional, como criar, abrir e gerenciar arquivos.

`io`: Oferece interfaces e utilitários para trabalhar com fluxos de entrada e saída (input/output), complementando as operações do os.

A combinação desses pacotes permite não apenas ler e escrever arquivos, mas também manipular permissões, diretórios e metadados. Vamos explorar cada detalhe passo a passo!

Exemplo 1: Veja este primeiro exemplo que mostrar como podemos criar e escrever em um arquivo em Go.

```go
arquivo, err := os.Create("exemplo.txt")

if err != nil {
    fmt.Println("Erro ao criar o arquivo:", err)
    return
}

defer arquivo.close()

// Escreve conteúdo no arquivo
conteudo := "Olá, este é um exemplo de texto!\n"

bytesEscritos, err := arquivo.WriteString(conteudo)

if err != nil {
    fmt.Println("Erro ao escrever no arquivo:", err)
    return
}

fmt.Printf("Arquivo criado com sucesso! %d bytes escritos.\n", bytesEscritos)
```