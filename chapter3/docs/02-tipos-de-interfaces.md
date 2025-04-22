# Tipos de Interfaces em Go

Go suporta diferentes tipos e estilos de interfaces, cada um com seus próprios casos de uso e benefícios.

## 1. Interfaces Pequenas

Go favorece interfaces pequenas e focadas. Este é um dos princípios fundamentais do design de interfaces em Go.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### Benefícios
- Mais fáceis de implementar
- Mais flexíveis
- Mais reutilizáveis
- Seguem o Princípio da Responsabilidade Única

## 2. Interfaces Compostas

Interfaces podem ser compostas de outras interfaces, permitindo criar interfaces mais complexas a partir de interfaces mais simples.

```go
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

### Exemplos de Composição
```go
type Closer interface {
    Close() error
}

type ReadCloser interface {
    Reader
    Closer
}

type WriteCloser interface {
    Writer
    Closer
}
```

## 3. Interfaces de Método Único

São muito comuns em Go e seguem o princípio de responsabilidade única.

```go
type Stringer interface {
    String() string
}

type Error interface {
    Error() string
}
```

## 4. Interfaces Comportamentais

Definem comportamentos específicos que tipos podem implementar.

```go
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

type Comparer interface {
    Compare(other interface{}) int
}
```

## 5. Interfaces Vazias

A interface vazia não define métodos e pode armazenar valores de qualquer tipo.

```go
interface{} // ou any no Go 1.18+
```

### Uso da Interface Vazia
```go
func PrintAny(v interface{}) {
    fmt.Printf("Valor: %v, Tipo: %T\n", v, v)
}
```

## Boas Práticas

1. **Mantenha Interfaces Pequenas**
   - Prefira interfaces com poucos métodos
   - Combine interfaces quando necessário

2. **Design por Composição**
   - Construa interfaces complexas combinando interfaces simples
   - Reutilize interfaces existentes

3. **Interfaces Focadas**
   - Cada interface deve ter um propósito claro
   - Evite interfaces que fazem muitas coisas diferentes

4. **Nomeação Clara**
   - Use nomes que descrevem o comportamento
   - Sufixo 'er' para interfaces de método único (Reader, Writer, etc.) 