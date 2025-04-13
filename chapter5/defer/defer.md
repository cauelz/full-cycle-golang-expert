# O Statement `defer` em Go

O `defer` em Go é uma palavra-chave usada para **adiar a execução de uma função ou método** até que a função que a contém termine sua execução (seja normalmente ou por um `panic`). Ele é comumente utilizado para garantir que recursos sejam liberados ou limpos, como fechar arquivos, conexões de rede, ou liberar locks, mesmo em casos de erros.

---

## Principais Usos do `defer`

### 1. **Gerenciamento de Recursos**
Garante que recursos alocados (como arquivos ou conexões) sejam liberados **automaticamente** após o uso, evitando vazamentos.

**Exemplo:**
```go
func readFile(filename string) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close() // Fecha o arquivo quando a função terminar

    // Processa o arquivo...
}
```

### 2. **Execução de Código no Final de uma Função**
Útil para ações que precisam ocorrer no final da função, como logs ou finalizações.

**Exemplo:**
```go
func exemplo() {
    fmt.Println("Início da função")
    defer fmt.Println("Executado ao final") // Será impresso por último
    fmt.Println("Fim da função")
}
// Saída:
// Início da função
// Fim da função
// Executado ao final
```

### 3. **Recuperação de `panic`**
Em conjunto com `recover`, o `defer` pode capturar e tratar `panics`, evitando a quebra do programa.

**Exemplo:**
```go
func safeOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do panic:", r)
        }
    }()
    // Código que pode causar panic...
}
```

---

## Comportamento do `defer`

### Ordem de Execução
Os statements `defer` são empilhados e executados em **ordem inversa** (LIFO: Last In, First Out).

**Exemplo:**
```go
func main() {
    defer fmt.Println("Primeiro defer")
    defer fmt.Println("Segundo defer")
    defer fmt.Println("Terceiro defer")
}
// Saída:
// Terceiro defer
// Segundo defer
// Primeiro defer
```

### Argumentos Avaliados Imediatamente
Os argumentos passados para `defer` são avaliados no momento da declaração, não na execução.

**Exemplo:**
```go
func exemplo() {
    x := 10
    defer fmt.Println("Valor de x:", x) // x = 10 aqui
    x = 20
}
// Saída: Valor de x: 10 (não 20)
```

---

## Interação com Valores de Retorno
Se a função retorna valores nomeados, um `defer` pode modificar esses valores antes do retorno.

**Exemplo:**
```go
func soma(a, b int) (resultado int) {
    defer func() {
        resultado += 10 // Modifica o valor de retorno
    }()
    return a + b // Retorno original: a + b
}

fmt.Println(soma(2, 3)) // Saída: 15 (2+3=5 + 10)
```

---

## Casos Comuns de Uso
- Fechar recursos (arquivos, conexões de rede, locks).
- Registrar eventos de finalização (logs, métricas).
- Tratamento de `panic` com `recover`.

---

## Armadilhas Comuns
1. **Loops com `defer`:** Usar `defer` dentro de loops pode causar acúmulo de chamadas (use funções anônimas para escopo limitado).
2. **Desempenho:** `defer` tem custo mínimo, mas em código crítico (ex.: loops intensivos), prefira gerenciamento manual.

---

# Resumo
O `defer` é uma ferramenta poderosa em Go para garantir **cleanup confiável** e **código mais legível**, especialmente em funções com múltiplos pontos de retorno ou possíveis erros. Use-o para centralizar ações pós-execução e evitar repetição de código.
