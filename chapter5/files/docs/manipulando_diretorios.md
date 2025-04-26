# Manipulando Diretórios em Go

Além de arquivos, o pacote `os` permite criar, remover e listar diretórios facilmente.

## Criando Diretórios

```go
err := os.Mkdir("meu_diretorio", 0755)
if err != nil {
    fmt.Println("Erro ao criar diretório:", err)
    return
}
```
- O segundo parâmetro define as permissões (estilo UNIX). Para detalhes sobre permissões de diretórios e exemplos práticos, consulte o arquivo [permissions.md](./permissions.md).

## Criando Diretórios Recursivamente

```go
err := os.MkdirAll("pasta1/pasta2/pasta3", 0755)
if err != nil {
    fmt.Println("Erro ao criar diretórios:", err)
    return
}
```

## Removendo Diretórios

```go
err := os.Remove("meu_diretorio")
if err != nil {
    fmt.Println("Erro ao remover diretório:", err)
}
```
- Só remove se estiver vazio.

## Removendo Diretórios Recursivamente

```go
err := os.RemoveAll("pasta1")
if err != nil {
    fmt.Println("Erro ao remover diretórios:", err)
}
```

## Listando Arquivos e Diretórios

```go
arquivos, err := os.ReadDir(".")
if err != nil {
    fmt.Println("Erro ao listar diretórios:", err)
    return
}
for _, arquivo := range arquivos {
    fmt.Println(arquivo.Name())
}
```

Essas funções facilitam a organização e manipulação de estruturas de pastas em Go. 