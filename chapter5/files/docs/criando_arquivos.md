# Criando Arquivos em Go

Para criar arquivos em Go, utilizamos a função `os.Create`. Ela cria um novo arquivo ou sobrescreve um existente.

## Exemplo Básico

```go
arquivo, err := os.Create("exemplo.txt")
if err != nil {
    fmt.Println("Erro ao criar o arquivo:", err)
    return
}
defer arquivo.Close()

conteudo := "Olá, este é um exemplo de texto!\n"
bytesEscritos, err := arquivo.WriteString(conteudo)
if err != nil {
    fmt.Println("Erro ao escrever no arquivo:", err)
    return
}
fmt.Printf("Arquivo criado com sucesso! %d bytes escritos.\n", bytesEscritos)
```

## Detalhes da Função
- `os.Create(nome)`: Cria ou sobrescreve o arquivo.
- Retorna: `*os.File` (ponteiro para o arquivo) e `error`.
- Sempre feche o arquivo com `defer arquivo.Close()` para evitar vazamento de recursos.

## Observações
- Se o arquivo já existir, seu conteúdo será apagado.
- Para evitar sobrescrita, utilize `os.OpenFile` com as flags apropriadas.
- Arquivos criados podem ter permissões definidas (por padrão, 0666 ajustado pelo umask do sistema). Para saber mais sobre permissões e como configurá-las corretamente, consulte o arquivo [permissions.md](./permissions.md). 