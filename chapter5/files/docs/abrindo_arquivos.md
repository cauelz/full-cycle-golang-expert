# Abrindo Arquivos em Go

Para ler ou manipular arquivos já existentes, usamos as funções `os.Open` e `os.OpenFile`.

## Usando os.Open

Abre um arquivo apenas para leitura.

```go
arquivo, err := os.Open("exemplo.txt")
if err != nil {
    fmt.Println("Erro ao abrir o arquivo:", err)
    return
}
defer arquivo.Close()
```

- Retorna: `*os.File` e `error`.
- Não cria o arquivo se não existir.

## Usando os.OpenFile

Permite abrir arquivos com mais controle (leitura, escrita, criação, append, etc).

```go
arquivo, err := os.OpenFile("exemplo.txt", os.O_RDWR|os.O_CREATE, 0666)
if err != nil {
    fmt.Println("Erro ao abrir/criar o arquivo:", err)
    return
}
defer arquivo.Close()
```

- Flags comuns:
  - `os.O_RDONLY`: Apenas leitura
  - `os.O_WRONLY`: Apenas escrita
  - `os.O_RDWR`: Leitura e escrita
  - `os.O_CREATE`: Cria se não existir
  - `os.O_APPEND`: Escreve no final
  - `os.O_TRUNC`: Limpa o conteúdo ao abrir
- Permissões: estilo UNIX (ex: 0664). Para entender como funcionam as permissões e como defini-las corretamente, consulte o arquivo [permissions.md](./permissions.md).

Consulte sempre a documentação para mais detalhes sobre as flags. 