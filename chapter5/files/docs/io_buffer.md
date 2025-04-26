# Leitura e Escrita Eficiente com bufio em Go

Para processar arquivos grandes, o pacote `bufio` oferece buffers que tornam a leitura e escrita mais eficiente.

## Lendo com bufio.NewReader

```go
arquivo, err := os.Open("grande.txt")
if err != nil {
    log.Fatal(err)
}
defer arquivo.Close()

leitor := bufio.NewReader(arquivo)
linha, err := leitor.ReadString('\n')
if err != nil && err != io.EOF {
    log.Fatal(err)
}
fmt.Print(linha)
```

## Escrevendo com bufio.NewWriter

```go
arquivo, err := os.Create("saida.txt")
if err != nil {
    log.Fatal(err)
}
defer arquivo.Close()

escritor := bufio.NewWriter(arquivo)
_, err = escritor.WriteString("Linha de exemplo\n")
if err != nil {
    log.Fatal(err)
}
escritor.Flush() // Importante para garantir que tudo seja gravado!
```

O uso de buffers é recomendado para arquivos grandes ou operações frequentes de leitura/escrita. 