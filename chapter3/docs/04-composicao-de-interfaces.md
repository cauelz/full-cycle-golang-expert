# Composição de Interfaces

A composição de interfaces é um recurso poderoso em Go que permite criar interfaces mais complexas combinando interfaces mais simples.

## Conceito Básico

Em Go, uma interface pode incluir outras interfaces em sua definição, herdando todos os seus métodos.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter combina as interfaces Reader e Writer
type ReadWriter interface {
    Reader
    Writer
}
```

## Benefícios da Composição

1. **Reutilização de Código**
   - Evita duplicação de definições
   - Mantém o código DRY (Don't Repeat Yourself)

2. **Modularidade**
   - Interfaces pequenas e focadas
   - Fácil composição para casos mais complexos

3. **Flexibilidade**
   - Adapta-se a diferentes necessidades
   - Permite evolução gradual

## Exemplos Práticos

### 1. Interface de IO
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Combinações comuns
type ReadWriter interface {
    Reader
    Writer
}

type ReadCloser interface {
    Reader
    Closer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

### 2. Interface de Gerenciamento
```go
type Identifier interface {
    ID() string
}

type Namer interface {
    Name() string
}

type Entity interface {
    Identifier
    Namer
    String() string
}
```

### 3. Interface de Manipulação de Dados
```go
type Loader interface {
    Load() error
}

type Saver interface {
    Save() error
}

type Processor interface {
    Process() error
}

type DataHandler interface {
    Loader
    Saver
    Processor
}
```

## Padrões de Composição

### 1. Composição Hierárquica
```go
type Base interface {
    Basic() error
}

type Advanced interface {
    Base
    Extra() bool
}

type Complete interface {
    Advanced
    Final() string
}
```

### 2. Composição Funcional
```go
type Calculator interface {
    Calculate() float64
}

type Validator interface {
    Validate() bool
}

type ProcessorWithValidation interface {
    Calculator
    Validator
    Process() error
}
```

## Boas Práticas

1. **Mantenha Interfaces Simples**
   ```go
   // Bom
   type Reader interface {
       Read(p []byte) (n int, err error)
   }

   // Evite interfaces muito grandes
   type DoEverything interface {
       // Muitos métodos...
   }
   ```

2. **Composição Lógica**
   ```go
   // Agrupe interfaces relacionadas
   type FileHandler interface {
       Reader
       Writer
       Closer
   }
   ```

3. **Nomeação Clara**
   ```go
   // Nome descreve a combinação
   type ReadWriter interface {
       Reader
       Writer
   }
   ```

## Casos de Uso Comuns

1. **Manipulação de Arquivos**
```go
type File interface {
    Reader
    Writer
    Closer
    Seeker
}
```

2. **APIs HTTP**
```go
type HTTPHandler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type MiddlewareHandler interface {
    HTTPHandler
    Middleware() Handler
}
```

3. **Banco de Dados**
```go
type Repository interface {
    Reader
    Writer
    Transactional
}

type Transactional interface {
    Begin() error
    Commit() error
    Rollback() error
}
```

## Dicas de Implementação

1. **Verificação de Implementação**
```go
// Verifica se *MyType implementa todas as interfaces
var _ ReadWriter = (*MyType)(nil)
```

2. **Documentação**
```go
// DataProcessor combina operações de leitura e processamento
type DataProcessor interface {
    Reader
    Processor
}
```

3. **Testes**
```go
// Mock que implementa múltiplas interfaces
type MockHandler struct{}

func (m *MockHandler) Read(p []byte) (int, error)  { return 0, nil }
func (m *MockHandler) Write(p []byte) (int, error) { return 0, nil }
func (m *MockHandler) Close() error                { return nil }
``` 