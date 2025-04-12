# Manipulação de Arquivos e Diretórios em Go.

Sempre tive muita dificuldade para entender como funcionam as operações com arquivos em linguagens de programação. Por isso, neste capítulo, vou me esforçar para que cada explicação seja didática, ajudando você a dominar esses conceitos de forma simples.

Podemos resumir a essência deste capítulo em dois verbos fundamentais: ler e escrever. Essas duas ações são a base de qualquer interação com arquivos e diretórios. Para realizá-las em Go, utilizamos principalmente dois pacotes da biblioteca padrão:

`os`: Responsável por operações de baixo nível com o sistema operacional, como criar, abrir e gerenciar arquivos.

`io`: Oferece interfaces e utilitários para trabalhar com fluxos de entrada e saída (input/output), complementando as operações do os.

A combinação desses pacotes permite não apenas ler e escrever arquivos, mas também manipular permissões, diretórios e metadados. Vamos explorar cada detalhe passo a passo!

Ainda neste capítulo, vamos falar sobre o pacote `buffer` que nos ajuda a processar arquivos grandes.

Exemplo 1: Veja este primeiro exemplo que mostra como podemos criar e escrever em um arquivo.

```go
// Para criar, podemos utilizar a função Create do pacote "os"

arquivo, err := os.Create("exemplo.txt")

if err != nil {
    fmt.Println("Erro ao criar o arquivo:", err)
    return
}

defer arquivo.close() // Fecha o arquivo ao finalizar o programa.

// Podemos escrever um texto neste arquivo criado com a função WriteString disponível a partir do arquivo criado.

conteudo := "Olá, este é um exemplo de texto!\n"

bytesEscritos, err := arquivo.WriteString(conteudo)

if err != nil {
    fmt.Println("Erro ao escrever no arquivo:", err)
    return
}

fmt.Printf("Arquivo criado com sucesso! %d bytes escritos.\n", bytesEscritos)
```

Vamos destrinchar? Veja como é fácil.

### `os.Create`  
- **O que faz?**  
  Cria um arquivo chamado `"exemplo.txt"` no diretório atual (ou **sobrescreve** se ele já existir!).  

- **Retornos:**  
  - `*os.File`: Ponteiro para manipular o arquivo.  
  - `error`: Retorna erro se houver problemas (ex.: permissão negada).  

---

### `defer arquivo.Close()`  
- **Funcionalidade:**  
  Garante que o arquivo será fechado **automaticamente** quando a função onde ele foi criado terminar.  

- **Por que usar?**  
  Evita **vazamento de recursos** (arquivos abertos consomem memória!).  

---

### `WriteString`  
- **O que faz?**  
  Escreve uma `string` no arquivo aberto.  

- **Retornos:**  
  - `int`: Número de bytes escritos.  
  - `error`: Possível erro durante a escrita.  

---

### ⚠️ **Cuidado!**  
Se `"exemplo.txt"` já existir:  
- `os.Create` **apaga todo o conteúdo anterior** sem aviso!  
- Para evitar isso, use `os.OpenFile` com a flag `os.O_APPEND`.

Agora que você teve uma breve introdução sobre o assunto, vamos explorar os pacotes?

Vamos o assunto da seguinte maneira:

- Manipulação de Arquivos
- Manipulação de Diretórios
- Acessar e modificar variáveis de ambiente
- Interagir com Input e Output (I/O)

---

# Manipulando Arquivos - Pacote OS (Operating System).

A linguagem Go oferece diversas funções no pacote `os` para manipular arquivos de forma simples e eficiente. Vamos explorar as principais funções?

## `os.Create()``

Responsável por `criar` ou `truncar`(apaga o conteúdo interno) de um arquivo existente.

`Sintaxe:`

```Go
func Create(name string) (*os.File, error)
```

`Parâmetros:`

- name: caminho (absoluto ou relativo) para o arquivo que se deseja criar.

`Retorno:`

- `*os.File`: um ponteiro para o objeto que representa o arquivo criado.
- `error`: retorna um erro durante a criação do arquivo. Os mais comuns são de `permissão do usuário` e `caminho inválido`.


## `os.Open()`

Responsável por `abrir` um arquivo no formato leitura.

`Sintaxe:`

```Go
func Open(name string) (*os.File, error)
```

`Parâmetros:`

- name: caminho para o arquivo ser aberto.

`Retorno:`

- `*os.File`: um ponteiro para o objeto que representa o arquivo aberto (modo leitura).
- `error`: retorna um erro durante a criação do arquivo.

## `os.OpenFile`

Responsável por `abrir` um arquivo, mas que nos permite ter controle sobre as permissões e flags. 

`Sintaxe:`

```Go
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
```

`Parâmetros:`

- name: caminho para o arquivo ser aberto.
- flag: um valor inteiro armazenado no pacote `os` que determina o formato de abertura do arquivo, ou seja, quais operações podemos fazer com o arquivo aberto.
    - os.O_RDONLYE: Apenas leitura.
    - os.O_WRONLY: Apenas escrita.
    - os.O_CREATE: Cria o arquivo se não existir.
    - os.O_APPEND: Escreve no final do arquivo.
    - os.O_TRUNC: Trunca o arquivo ao abrir.
- perm: Define as permissões do arquivo(estilo UNIX). Exemplo: 0664 permite leitura e escrita para o dono e somente leitura para o grupo e outros. [Guia de Permissões em UNIX](https://github.com/cauelz/full-cycle-golang-expert/blob/master/chapter5/permissions/permissions.md)

`Retorno:`

- `*os.File`: um ponteiro para o objeto que representa o arquivo aberto (modo leitura).
- `error`: retorna um erro durante a criação do arquivo.