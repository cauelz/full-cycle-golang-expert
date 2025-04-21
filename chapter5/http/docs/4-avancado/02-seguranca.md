# Segurança e TLS em Go

## Introdução

A segurança é um aspecto crítico em aplicações web. O Go oferece suporte robusto para HTTPS/TLS e várias práticas de segurança importantes. Vamos explorar como implementar essas medidas de segurança.

## 1. HTTPS/TLS

### 1.1 Servidor HTTPS Básico

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    server := &http.Server{
        Addr:    ":443",
        Handler: mux,
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    }

    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
```

### 1.2 Gerando Certificados para Desenvolvimento

```bash
# Gerar chave privada
openssl genrsa -out key.pem 2048

# Gerar certificado auto-assinado
openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 365
```

### 1.3 Configuração TLS Customizada

```go
func configureTLS() *tls.Config {
    return &tls.Config{
        MinVersion: tls.VersionTLS12,
        CurvePreferences: []tls.CurveID{
            tls.X25519,
            tls.CurveP256,
        },
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
        },
    }
}
```

## 2. Headers de Segurança

### 2.1 Middleware de Headers de Segurança

```go
func securityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Prevenir clickjacking
        w.Header().Set("X-Frame-Options", "DENY")
        
        // Habilitar proteção XSS no browser
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        
        // Prevenir MIME sniffing
        w.Header().Set("X-Content-Type-Options", "nosniff")
        
        // Política de segurança de conteúdo
        w.Header().Set("Content-Security-Policy", 
            "default-src 'self'; script-src 'self'")
        
        // HSTS (só em produção com HTTPS)
        w.Header().Set("Strict-Transport-Security",
            "max-age=31536000; includeSubDomains")

        next.ServeHTTP(w, r)
    })
}
```

## 3. Autenticação

### 3.1 JWT (JSON Web Tokens)

```go
type Claims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func createToken(userID, role string, secret []byte) (string, error) {
    claims := Claims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

func validateToken(tokenString string, secret []byte) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, 
        func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("método de assinatura inesperado: %v", 
                    token.Header["alg"])
            }
            return secret, nil
        })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("token inválido")
}
```

### 3.2 Middleware de Autenticação

```go
func authMiddleware(secret []byte) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Extrair token do header
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", 
                    http.StatusUnauthorized)
                return
            }

            // Formato: "Bearer <token>"
            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, "Invalid authorization format", 
                    http.StatusUnauthorized)
                return
            }

            // Validar token
            claims, err := validateToken(parts[1], secret)
            if err != nil {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            // Adicionar claims ao contexto
            ctx := context.WithValue(r.Context(), "claims", claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

## 4. Proteção contra Ataques Comuns

### 4.1 CSRF (Cross-Site Request Forgery)

```go
type CSRFToken struct {
    Secret []byte
    Cookie string
}

func (c *CSRFToken) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            // Gerar token para GETs
            token := generateToken(c.Secret)
            http.SetCookie(w, &http.Cookie{
                Name:     c.Cookie,
                Value:    token,
                HttpOnly: true,
                Secure:   true,
                SameSite: http.SameSiteStrictMode,
            })
        } else {
            // Validar token para outros métodos
            cookie, err := r.Cookie(c.Cookie)
            if err != nil {
                http.Error(w, "CSRF cookie not found", 
                    http.StatusBadRequest)
                return
            }

            token := r.Header.Get("X-CSRF-Token")
            if !validateCSRFToken(token, cookie.Value, c.Secret) {
                http.Error(w, "Invalid CSRF token", 
                    http.StatusForbidden)
                return
            }
        }

        next.ServeHTTP(w, r)
    })
}
```

### 4.2 Rate Limiting por IP

```go
type IPRateLimiter struct {
    ips    map[string]*rate.Limiter
    mu     sync.RWMutex
    rate   rate.Limit
    burst  int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
    return &IPRateLimiter{
        ips:   make(map[string]*rate.Limiter),
        rate:  r,
        burst: b,
    }
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
    i.mu.Lock()
    defer i.mu.Unlock()

    limiter, exists := i.ips[ip]
    if !exists {
        limiter = rate.NewLimiter(i.rate, i.burst)
        i.ips[ip] = limiter
    }

    return limiter
}

func (i *IPRateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := getIP(r)
        limiter := i.GetLimiter(ip)
        
        if !limiter.Allow() {
            http.Error(w, "Too many requests", 
                http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

## 5. Validação de Entrada

### 5.1 Sanitização de Dados

```go
func sanitizeInput(input string) string {
    // Remover caracteres perigosos
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsNumber(r) {
            return r
        }
        return -1
    }, input)
}

// Uso em handlers
func handleUserInput(w http.ResponseWriter, r *http.Request) {
    input := r.FormValue("user_input")
    sanitized := sanitizeInput(input)
    
    // Usar input sanitizado
    processInput(sanitized)
}
```

### 5.2 Validação de JSON

```go
type User struct {
    Username string `json:"username" validate:"required,min=3,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,gte=18,lte=120"`
}

func validateJSON(data []byte, schema interface{}) error {
    validate := validator.New()
    
    if err := json.Unmarshal(data, schema); err != nil {
        return fmt.Errorf("erro ao decodificar JSON: %v", err)
    }
    
    if err := validate.Struct(schema); err != nil {
        return fmt.Errorf("erro de validação: %v", err)
    }
    
    return nil
}
```

## 6. Logging Seguro

```go
func secureLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Remover dados sensíveis dos logs
        cleanURL := sanitizeURL(r.URL.String())
        cleanHeaders := sanitizeHeaders(r.Header)

        // Log seguro
        log.Printf("Request: method=%s url=%s headers=%v",
            r.Method, cleanURL, cleanHeaders)

        next.ServeHTTP(w, r)
    })
}

func sanitizeURL(url string) string {
    // Remover parâmetros sensíveis
    sensitive := []string{"password", "token", "key"}
    for _, s := range sensitive {
        re := regexp.MustCompile(`(?i)`+s+`=[^&]*`)
        url = re.ReplaceAllString(url, s+"=REDACTED")
    }
    return url
}
```

## 7. Boas Práticas

1. **Sempre use HTTPS em produção**
   ```go
   // Redirecionar HTTP para HTTPS
   go http.ListenAndServe(":80", http.HandlerFunc(redirect))
   log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux))
   ```

2. **Rotação de Segredos**
   ```go
   type KeyRotator struct {
       current []byte
       previous []byte
       mu sync.RWMutex
   }

   func (kr *KeyRotator) Rotate(newKey []byte) {
       kr.mu.Lock()
       defer kr.mu.Unlock()
       kr.previous = kr.current
       kr.current = newKey
   }
   ```

3. **Timeouts Apropriados**
   ```go
   server := &http.Server{
       ReadTimeout:  5 * time.Second,
       WriteTimeout: 10 * time.Second,
       IdleTimeout:  120 * time.Second,
   }
   ```

4. **Limitar Tamanho do Payload**
   ```go
   func maxBodyBytes(h http.Handler, n int64) http.Handler {
       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           r.Body = http.MaxBytesReader(w, r.Body, n)
           h.ServeHTTP(w, r)
       })
   }
   ```

## Próximos Passos

- [Performance e Otimização](03-performance.md)
- [Testes de Segurança](04-testes-seguranca.md)
- [Monitoramento](05-monitoramento.md) 