# Cliente HTTP em Go

## Introdução

O pacote `net/http` em Go oferece um cliente HTTP robusto e fácil de usar. Vamos explorar desde os usos mais básicos até configurações avançadas e boas práticas.

## 1. Cliente Básico

### 1.1 Funções Helper

```go
// GET simples
resp, err := http.Get("https://api.exemplo.com/users")

// POST simples
resp, err := http.Post("https://api.exemplo.com/users", 
    "application/json", 
    strings.NewReader(`{"name": "João"}`))

// POST form
resp, err := http.PostForm("https://api.exemplo.com/form", 
    url.Values{"key": {"value"}})
```

### 1.2 Usando o Response

```go
resp, err := http.Get("https://api.exemplo.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close() // Sempre feche o body!

// Ler o corpo
body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}

// Decodificar JSON
var users []User
if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
    log.Fatal(err)
}
```

## 2. Cliente Customizado

### 2.1 Criando um Cliente

```go
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 100,
        IdleConnTimeout:     90 * time.Second,
        TLSHandshakeTimeout: 10 * time.Second,
    },
}
```

### 2.2 Usando o Cliente

```go
req, err := http.NewRequest("GET", "https://api.exemplo.com/users", nil)
if err != nil {
    log.Fatal(err)
}

// Adicionar headers
req.Header.Set("Authorization", "Bearer " + token)
req.Header.Set("Accept", "application/json")

// Fazer a requisição
resp, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

## 3. Contexto e Cancelamento

### 3.1 Timeout por Requisição

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, "GET", 
    "https://api.exemplo.com/users", nil)
if err != nil {
    log.Fatal(err)
}

resp, err := client.Do(req)
```

### 3.2 Cancelamento Manual

```go
ctx, cancel := context.WithCancel(context.Background())

// Em outra goroutine
go func() {
    time.Sleep(2 * time.Second)
    cancel() // Cancela a requisição após 2 segundos
}()

req, _ := http.NewRequestWithContext(ctx, "GET", 
    "https://api.exemplo.com/users", nil)
resp, err := client.Do(req)
```

## 4. Enviando Dados

### 4.1 POST com JSON

```go
user := User{Name: "João", Email: "joao@exemplo.com"}
data, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}

req, err := http.NewRequest("POST", "https://api.exemplo.com/users",
    bytes.NewBuffer(data))
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
```

### 4.2 Multipart Form

```go
// Criar um buffer para o form
var buf bytes.Buffer
writer := multipart.NewWriter(&buf)

// Adicionar campos
writer.WriteField("name", "João")

// Adicionar arquivo
file, _ := os.Open("foto.jpg")
defer file.Close()
part, _ := writer.CreateFormFile("photo", "foto.jpg")
io.Copy(part, file)

writer.Close()

// Criar requisição
req, err := http.NewRequest("POST", "https://api.exemplo.com/upload",
    &buf)
req.Header.Set("Content-Type", writer.FormDataContentType())

resp, err := client.Do(req)
```

## 5. Tratamento de Erros

### 5.1 Verificando Status Code

```go
resp, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
    body, _ := io.ReadAll(resp.Body)
    log.Fatalf("status: %d, body: %s", resp.StatusCode, body)
}
```

### 5.2 Retry com Backoff

```go
func doWithRetry(client *http.Client, req *http.Request) (*http.Response, error) {
    maxRetries := 3
    backoff := 100 * time.Millisecond

    var resp *http.Response
    var err error

    for i := 0; i < maxRetries; i++ {
        resp, err = client.Do(req)
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }
        
        if resp != nil {
            resp.Body.Close()
        }

        // Esperar antes de tentar novamente
        time.Sleep(backoff)
        backoff *= 2 // Backoff exponencial
    }

    return resp, err
}
```

## 6. Boas Práticas

1. **Reutilize Clientes**
   ```go
   // Ruim: criar cliente por requisição
   resp, err := (&http.Client{}).Get(url)

   // Bom: reutilizar cliente
   var client = &http.Client{} // global ou injetado
   resp, err := client.Get(url)
   ```

2. **Configure Timeouts**
   ```go
   client := &http.Client{
       Timeout: 10 * time.Second,
   }
   ```

3. **Limite Conexões**
   ```go
   transport := &http.Transport{
       MaxIdleConns:        100,
       MaxIdleConnsPerHost: 100,
       MaxConnsPerHost:     100,
   }
   ```

4. **Use Contexto**
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
   defer cancel()
   req = req.WithContext(ctx)
   ```

5. **Sempre Feche o Body**
   ```go
   resp, err := client.Do(req)
   if err != nil {
       return err
   }
   defer resp.Body.Close()
   ```

## 7. Exemplo Completo

```go
type APIClient struct {
    client  *http.Client
    baseURL string
    token   string
}

func NewAPIClient(baseURL, token string) *APIClient {
    return &APIClient{
        client: &http.Client{
            Timeout: 10 * time.Second,
            Transport: &http.Transport{
                MaxIdleConns:        100,
                MaxIdleConnsPerHost: 100,
                IdleConnTimeout:     90 * time.Second,
            },
        },
        baseURL: baseURL,
        token:   token,
    }
}

func (c *APIClient) GetUser(ctx context.Context, id string) (*User, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", 
        fmt.Sprintf("%s/users/%s", c.baseURL, id), nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", "Bearer "+c.token)
    req.Header.Set("Accept", "application/json")

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("status: %d, body: %s", 
            resp.StatusCode, body)
    }

    var user User
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}
```

## Próximos Passos

- [Autenticação e Autorização](02-autenticacao.md)
- [Middlewares de Cliente](03-middlewares-cliente.md)
- [Testes de Cliente HTTP](04-testes.md) 