# Exercício 03: Adicionar Texto ao Final de um Arquivo (Append)

Escreva um programa em Go que:

1. Abra o arquivo `mensagem.txt` para escrita no modo append.
2. Adicione a linha "Esta é uma nova linha adicionada ao arquivo." ao final do arquivo.
3. Feche o arquivo corretamente.

Dicas:
- Use o modo `os.O_APPEND` ao abrir o arquivo.
- Trate erros de abertura e escrita. 

---

## Observação Importante sobre os Flags de Abertura

> **Atenção:** Para abrir um arquivo para escrita no modo append, é necessário combinar o flag `os.O_APPEND` com `os.O_WRONLY` (ou `os.O_RDWR`). Apenas `os.O_APPEND` não é suficiente, pois ele apenas indica que a escrita será feita ao final do arquivo, mas não habilita a escrita em si. Exemplo:
>
> ```go
> file, err := os.OpenFile("mensagem.txt", os.O_APPEND|os.O_WRONLY, 0644)
> ```

---

## Sobre Permissões ao Abrir Arquivos em Go

O terceiro argumento da função `os.OpenFile` define as permissões do arquivo, caso ele seja criado. Ele é representado em octal (ex: `0644`). Os valores mais comuns são:

- `0644`: Permite leitura e escrita para o dono, e apenas leitura para os outros.
- `0600`: Permite leitura e escrita apenas para o dono.
- `0666`: Permite leitura e escrita para todos.

Exemplo de uso:

```go
file, err := os.OpenFile("arquivo.txt", os.O_CREATE|os.O_WRONLY, 0644)
```

Essas permissões só têm efeito se o arquivo for criado. Se o arquivo já existir, as permissões não são alteradas. 