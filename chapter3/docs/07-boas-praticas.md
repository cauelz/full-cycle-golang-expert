# Boas Práticas com Interfaces em Go

Este documento apresenta um conjunto abrangente de boas práticas para o uso de interfaces em Go, ajudando a criar código mais limpo, manutenível e eficiente.

## Princípios Fundamentais

### 1. Interfaces Pequenas
```go
// Bom
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Evite interfaces grandes
type DoEverything interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
    Flush() error
    // ... muitos outros métodos
}
```

### 2. Interface Segregation Principle (ISP)
Prefira múltiplas interfaces pequenas a uma grande interface.

```go
// Bom
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Padrões de Design

### 1. Aceite Interfaces, Retorne Structs
```go
// Bom
func ProcessReader(r Reader) *Result {
    // ...
}

// Evite
func ProcessReader() Reader {
    // ...
}
```

### 2. Composição sobre Herança
```go
// Bom
type Logger struct {
    writer Writer
}

// Evite simular herança
type SpecialWriter struct {
    Writer  // embedding
}
```

## Implementação

### 1. Implementação Implícita
```go
// Definição da interface
type Stringer interface {
    String() string
}

// Implementação implícita
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s, %d anos", p.Name, p.Age)
}
```

### 2. Verificação de Implementação
```go
// Verifica em tempo de compilação
var _ json.Marshaler = (*CustomType)(nil)
var _ io.Reader = (*CustomReader)(nil)
```

## Documentação

### 1. Comentários de Interface
```go
// Reader é a interface que encapsula o método básico Read.
//
// Read lê até len(p) bytes no slice p e retorna o número de
// bytes lidos (0 <= n <= len(p)) e qualquer erro encontrado.
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### 2. Exemplos de Uso
```go
// Example_reader demonstra como usar a interface Reader
func Example_reader() {
    r := strings.NewReader("hello")
    buf := make([]byte, 5)
    n, _ := r.Read(buf)
    fmt.Printf("%d bytes: %s\n", n, buf)
    // Output: 5 bytes: hello
}
```

## Testes

### 1. Mock de Interfaces
```go
type MockReader struct {
    ReadFunc func(p []byte) (n int, err error)
}

func (m *MockReader) Read(p []byte) (n int, err error) {
    return m.ReadFunc(p)
}

func TestWithMock(t *testing.T) {
    mock := &MockReader{
        ReadFunc: func(p []byte) (n int, err error) {
            return len(p), nil
        },
    }
    // Use mock em testes
}
```

### 2. Interfaces para Testabilidade
```go
type TimeProvider interface {
    Now() time.Time
}

type RealTime struct{}

func (RealTime) Now() time.Time {
    return time.Now()
}

type MockTime struct {
    CurrentTime time.Time
}

func (m MockTime) Now() time.Time {
    return m.CurrentTime
}
```

## Tratamento de Erros

### 1. Interface error
```go
type CustomError struct {
    Code    int
    Message string
}

func (e CustomError) Error() string {
    return fmt.Sprintf("erro %d: %s", e.Code, e.Message)
}
```

### 2. Wrapping de Erros
```go
type ErrorWrapper interface {
    Error() string
    Unwrap() error
}

type WrapError struct {
    err error
    msg string
}

func (w *WrapError) Error() string {
    return fmt.Sprintf("%s: %v", w.msg, w.err)
}

func (w *WrapError) Unwrap() error {
    return w.err
}
```

## Performance

### 1. Evite Interface{} Desnecessária
```go
// Evite
func ProcessAny(data interface{}) {
    // ...
}

// Prefira
func ProcessString(data string) {
    // ...
}
```

### 2. Minimize Conversões de Tipo
```go
// Evite conversões frequentes
for i := 0; i < 1000; i++ {
    val := interface{}(i).(int)
    // ...
}

// Melhor
func process(i int) {
    // trabalhe diretamente com o tipo
}
```

## Convenções de Nomenclatura

### 1. Nomes de Interfaces
```go
// Interface de método único: use 'er'
type Reader interface { ... }
type Writer interface { ... }

// Interfaces comportamentais: use substantivos
type Cache interface { ... }
type Storage interface { ... }
```

### 2. Métodos de Interface
```go
// Use nomes concisos mas descritivos
type Processor interface {
    Process() error
    ProcessWith(opts Options) error
}
```

## Anti-Padrões

### 1. Evite Interfaces Vazias Desnecessárias
```go
// Evite
func ProcessAnything(v interface{}) {
    // ...
}

// Prefira tipos específicos quando possível
func ProcessStrings(v []string) {
    // ...
}
```

### 2. Não Force Implementações Desnecessárias
```go
// Evite interfaces muito específicas
type UserProcessor interface {
    ProcessUser(u User) error
    ValidateUser(u User) bool
    FormatUser(u User) string
    // ... muitos métodos específicos
}

// Prefira interfaces focadas
type UserValidator interface {
    ValidateUser(u User) bool
}
```

## Dicas Adicionais

1. **Mantenha a Coesão**
   - Interfaces devem ter um propósito claro
   - Métodos relacionados devem estar juntos

2. **Favoreça Composição**
   - Combine interfaces pequenas
   - Evite hierarquias profundas

3. **Documente Comportamentos**
   - Explique o propósito da interface
   - Documente casos especiais

4. **Pense na Evolução**
   - Interfaces são contratos
   - Mudanças podem afetar muitos consumidores 