# Conceitos Básicos de Interfaces

Uma interface em Go é um tipo abstrato que define um conjunto de métodos. Ela especifica o comportamento que um tipo deve ter, sem se preocupar com a implementação concreta.

## Definição Básica

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

## Características Principais

1. **Abstração de Comportamento**
   - Interfaces definem o "o que" sem especificar o "como"
   - Permitem múltiplas implementações do mesmo comportamento

2. **Contrato de Métodos**
   - Lista de métodos que um tipo deve implementar
   - Não inclui implementação, apenas assinaturas

3. **Flexibilidade**
   - Permite trocar implementações facilmente
   - Facilita testes e mocks

## Uso Básico

```go
// Definição da interface
type Printer interface {
    Print() string
}

// Implementação da interface
type Document struct {
    content string
}

// Document implementa Printer
func (d Document) Print() string {
    return d.content
}

// Função que usa a interface
func PrintDocument(p Printer) {
    fmt.Println(p.Print())
}
```

## Quando Usar

- Para definir comportamentos comuns entre diferentes tipos
- Quando precisar de flexibilidade na implementação
- Para facilitar testes unitários
- Quando quiser desacoplar componentes do sistema 