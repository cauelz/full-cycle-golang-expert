# Handlers e ServeMux em Go

## Introdução

Em Go, o sistema de roteamento HTTP é baseado em dois conceitos principais:
1. **Handlers**: Responsáveis por processar requisições HTTP
2. **ServeMux**: O roteador que direciona requisições para os handlers corretos

## 1. Handlers

### 1.1 A Interface Handler

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Esta interface é o coração do sistema de roteamento em Go. Qualquer tipo que implemente este método pode ser usado como um handler HTTP.

### 1.2 HandlerFunc

```go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

`HandlerFunc` é um adaptador que permite usar funções comuns como handlers HTTP.

### 1.3 Exemplos de Handlers

#### Handler Função


```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Olá, Mundo!")
}

// Uso:
http.HandleFunc("/hello", helloHandler)
```

#### Handler Struct
```go
type userHandler struct {
    users map[string]User
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        json.NewEncoder(w).Encode(h.users)
    case "POST":
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        // ... salvar usuário
    }
}

// Uso:
handler := &userHandler{users: make(map[string]User)}
http.Handle("/users", handler)
```

## 2. ServeMux

### 2.1 O que é ServeMux?

O conceito de multiplexador (mux) vem da eletrônica e telecomunicações. Um multiplexador é um dispositivo que seleciona uma entre várias entradas e encaminha a selecionada para uma única saída. Na eletrônica digital, multiplexadores são usados para:

- Combinar múltiplos sinais em um único canal
- Rotear diferentes entradas para uma saída específica
- Economizar recursos compartilhando um canal

No contexto HTTP, o ServeMux aplica o mesmo conceito:
- As "entradas" são as diferentes URLs/rotas
- A "saída" é o handler que vai processar a requisição
- O "seletor" é o padrão de rota que determina qual handler usar

ServeMux é o multiplexador HTTP padrão em Go. Ele:
- Mapeia URLs para handlers
- Suporta padrões de rota
- É seguro para uso concorrente
- Não cria novas goroutines

### 2.2 DefaultServeMux

```go
// Usando o ServeMux padrão global
http.HandleFunc("/", homeHandler)
http.Handle("/api/", apiHandler)
http.ListenAndServe(":8080", nil) // nil usa DefaultServeMux
```

### 2.3 ServeMux Customizado

```go
mux := http.NewServeMux()

// Registrar handlers
mux.HandleFunc("/", homeHandler)
mux.Handle("/api/", apiHandler)

// Usar o mux customizado
http.ListenAndServe(":8080", mux)
```

## 3. Padrões de Rota

### 3.1 Rota Exata
```go
mux.HandleFunc("/about", aboutHandler)
// Corresponde apenas a /about
```

### 3.2 Rota Prefixo
```go
mux.Handle("/static/", http.StripPrefix("/static/", 
    http.FileServer(http.Dir("static"))))
// Corresponde a /static/ e qualquer coisa abaixo
```

### 3.3 Novos Padrões (Go 1.22+)

```go
// Rota com método HTTP
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)

// Rota com variáveis
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("PUT /users/{id}", updateUser)

// Wildcard
mux.HandleFunc("GET /files/*", serveFiles)
```

## 4. Organização de Handlers

### 4.1 Por Domínio

```go
type UserHandler struct {
    store UserStore
}

func (h *UserHandler) Routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /users", h.List)
    mux.HandleFunc("POST /users", h.Create)
    mux.HandleFunc("GET /users/{id}", h.Get)
    mux.HandleFunc("PUT /users/{id}", h.Update)
    return mux
}

// Uso:
userHandler := &UserHandler{store: store}
mainMux.Handle("/users/", userHandler.Routes())
```

### 4.2 Por Funcionalidade

```go
func setupAuthRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /login", handleLogin)
    mux.HandleFunc("POST /logout", handleLogout)
    mux.HandleFunc("POST /refresh", handleRefreshToken)
}

func setupAPIRoutes(mux *http.ServeMux) {
    mux.Handle("/api/users/", userHandler)
    mux.Handle("/api/products/", productHandler)
    mux.Handle("/api/orders/", orderHandler)
}

func main() {
    mux := http.NewServeMux()
    setupAuthRoutes(mux)
    setupAPIRoutes(mux)
    http.ListenAndServe(":8080", mux)
}
```

## 5. Boas Práticas

### 5.1 Organização de Código

1. **Separe por Domínio**
   ```go
   /handlers
     /user
       list.go
       create.go
       update.go
     /product
       list.go
       create.go
   ```

2. **Use Interfaces**
   ```go
   type UserService interface {
       List() ([]User, error)
       Create(User) error
       Get(id string) (User, error)
   }

   type UserHandler struct {
       service UserService
   }
   ```

3. **Injeção de Dependências**
   ```go
   func NewUserHandler(service UserService) *UserHandler {
       return &UserHandler{service: service}
   }
   ```

### 5.2 Tratamento de Erros

```go
func (h *UserHandler) handleError(w http.ResponseWriter, err error) {
    switch e := err.(type) {
    case *NotFoundError:
        http.Error(w, e.Error(), http.StatusNotFound)
    case *ValidationError:
        http.Error(w, e.Error(), http.StatusBadRequest)
    default:
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.service.Get(id)
    if err != nil {
        h.handleError(w, err)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

### 5.3 Respostas JSON

```go
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
    users, err := h.service.List()
    if err != nil {
        h.handleError(w, err)
        return
    }
    respondJSON(w, http.StatusOK, users)
}
```

## 6. Exemplos Práticos

### 6.1 API REST Completa

```go
type API struct {
    users    *UserHandler
    products *ProductHandler
    auth     *AuthHandler
}

func NewAPI(db *sql.DB) *API {
    userService := NewUserService(db)
    productService := NewProductService(db)
    authService := NewAuthService(db)

    return &API{
        users:    NewUserHandler(userService),
        products: NewProductHandler(productService),
        auth:     NewAuthHandler(authService),
    }
}

func (api *API) Routes() *http.ServeMux {
    mux := http.NewServeMux()
    
    // Auth routes
    mux.HandleFunc("POST /login", api.auth.Login)
    mux.HandleFunc("POST /logout", api.auth.Logout)
    
    // API routes
    mux.Handle("/api/users/", api.users.Routes())
    mux.Handle("/api/products/", api.products.Routes())
    
    return mux
}
```

## Próximos Passos

- [Ciclo de Vida de uma Requisição](03-ciclo-de-vida.md)
- [Middlewares](../4-avancado/01-middlewares.md)
- [Performance e Otimização](../4-avancado/03-performance.md) 