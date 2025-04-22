# Type Assertions e Type Switches em Go

Type assertions e type switches são mecanismos fundamentais em Go para trabalhar com interfaces e determinar tipos em tempo de execução.

## Type Assertions

Uma type assertion fornece acesso ao tipo concreto de um valor de interface.

### Sintaxe Básica
```go
value, ok := interface.(Type)
```

### Exemplos

1. **Assertion Simples**
```go
var i interface{} = "hello"

// Com verificação de tipo segura
if str, ok := i.(string); ok {
    fmt.Println(str)
} else {
    fmt.Println("não é uma string")
}

// Assertion direta (pode causar panic)
str := i.(string)
```

2. **Múltiplas Assertions**
```go
func processValue(i interface{}) string {
    if str, ok := i.(string); ok {
        return str
    }
    if num, ok := i.(int); ok {
        return strconv.Itoa(num)
    }
    return "tipo desconhecido"
}
```

## Type Switches

Type switches permitem fazer múltiplas type assertions de forma mais elegante.

### Sintaxe Básica
```go
switch v := i.(type) {
case Type1:
    // code
case Type2:
    // code
default:
    // code
}
```

### Exemplos

1. **Switch Básico**
```go
func describe(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Printf("String com tamanho %v\n", len(v))
    case int:
        fmt.Printf("Integer com valor %v\n", v)
    case bool:
        fmt.Printf("Boolean com valor %v\n", v)
    default:
        fmt.Printf("Tipo desconhecido!\n")
    }
}
```

2. **Switch com Múltiplos Tipos**
```go
func categorize(i interface{}) string {
    switch v := i.(type) {
    case int, int32, int64:
        return "número inteiro"
    case float32, float64:
        return "número decimal"
    case string:
        return "texto"
    case nil:
        return "nulo"
    default:
        return fmt.Sprintf("tipo desconhecido: %T", v)
    }
}
```

## Boas Práticas

### 1. Sempre Use Verificação de Tipo
```go
// Bom
if val, ok := i.(string); ok {
    // use val com segurança
}

// Ruim - pode causar panic
val := i.(string)
```

### 2. Organize Type Switches
```go
func process(i interface{}) error {
    switch v := i.(type) {
    case string:
        return processString(v)
    case int:
        return processInt(v)
    case error:
        return v
    default:
        return fmt.Errorf("tipo não suportado: %T", v)
    }
}
```

### 3. Use Funções Auxiliares
```go
func isString(i interface{}) bool {
    _, ok := i.(string)
    return ok
}

func toString(i interface{}) (string, error) {
    s, ok := i.(string)
    if !ok {
        return "", fmt.Errorf("não é uma string: %v", i)
    }
    return s, nil
}
```

## Casos de Uso Comuns

### 1. Processamento de Dados Genéricos
```go
func process(data interface{}) error {
    switch v := data.(type) {
    case []byte:
        return processBytes(v)
    case string:
        return processString(v)
    case io.Reader:
        return processReader(v)
    default:
        return fmt.Errorf("tipo não suportado: %T", v)
    }
}
```

### 2. Validação de Tipos
```go
type Validator interface {
    Validate() bool
}

func validateAny(i interface{}) bool {
    if v, ok := i.(Validator); ok {
        return v.Validate()
    }
    return false
}
```

### 3. Conversão de Tipos
```go
func convert(i interface{}) string {
    switch v := i.(type) {
    case string:
        return v
    case int:
        return strconv.Itoa(v)
    case float64:
        return strconv.FormatFloat(v, 'f', -1, 64)
    case bool:
        return strconv.FormatBool(v)
    default:
        return fmt.Sprintf("%v", v)
    }
}
```

## Padrões Avançados

### 1. Type Assertion em Interfaces Específicas
```go
type Sizer interface {
    Size() int
}

func getSize(i interface{}) int {
    if s, ok := i.(Sizer); ok {
        return s.Size()
    }
    return 0
}
```

### 2. Type Switch com Fallthrough
```go
func classify(i interface{}) string {
    switch v := i.(type) {
    case nil:
        return "nulo"
    case int, int8, int16, int32, int64:
        return fmt.Sprintf("inteiro: %d", v)
    case float32, float64:
        return fmt.Sprintf("decimal: %f", v)
    default:
        return fmt.Sprintf("outro: %T", v)
    }
}
```

### 3. Type Assertion Encadeada
```go
func process(i interface{}) error {
    if str, ok := i.(string); ok {
        return processString(str)
    } else if num, ok := i.(int); ok {
        return processInt(num)
    } else if arr, ok := i.([]interface{}); ok {
        return processArray(arr)
    }
    return fmt.Errorf("tipo não suportado: %T", i)
}
```

## Dicas de Debug

### 1. Impressão de Tipo
```go
func debugType(i interface{}) {
    fmt.Printf("Tipo: %T, Valor: %v\n", i, i)
}
```

### 2. Verificação de Interface
```go
func debugInterface(i interface{}) {
    fmt.Printf("Tipo: %T\n", i)
    fmt.Printf("É error? %v\n", isError(i))
    fmt.Printf("É fmt.Stringer? %v\n", isStringer(i))
}

func isError(i interface{}) bool {
    _, ok := i.(error)
    return ok
}

func isStringer(i interface{}) bool {
    _, ok := i.(fmt.Stringer)
    return ok
}
``` 