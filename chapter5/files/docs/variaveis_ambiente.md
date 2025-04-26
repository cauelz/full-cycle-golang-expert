# Variáveis de Ambiente em Go

O pacote `os` permite acessar e modificar variáveis de ambiente do sistema operacional.

## Lendo Variáveis de Ambiente

```go
valor := os.Getenv("NOME_VARIAVEL")
fmt.Println("Valor:", valor)
```

## Definindo Variáveis de Ambiente

```go
err := os.Setenv("NOME_VARIAVEL", "valor")
if err != nil {
    fmt.Println("Erro ao definir variável:", err)
}
```

## Removendo Variáveis de Ambiente

```go
err := os.Unsetenv("NOME_VARIAVEL")
if err != nil {
    fmt.Println("Erro ao remover variável:", err)
}
```

Essas funções são úteis para configurar comportamentos do programa sem alterar o código-fonte. 