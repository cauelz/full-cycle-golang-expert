package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/time/rate"
)

// User representa um usuário do sistema
type User struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role"`
}

// UserStore é um armazenamento em memória para usuários
type UserStore struct {
	sync.RWMutex
	users map[string]User
}

// NewUserStore cria um novo armazenamento de usuários
func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]User),
	}
}

// Add adiciona um novo usuário
func (s *UserStore) Add(user User) {
	s.Lock()
	defer s.Unlock()
	s.users[user.ID] = user
}

// Get retorna um usuário pelo ID
func (s *UserStore) Get(id string) (User, bool) {
	s.RLock()
	defer s.RUnlock()
	user, ok := s.users[id]
	return user, ok
}

// GetByUsername retorna um usuário pelo username
func (s *UserStore) GetByUsername(username string) (User, bool) {
	s.RLock()
	defer s.RUnlock()
	for _, user := range s.users {
		if user.Username == username {
			return user, true
		}
	}
	return User{}, false
}

// Claims representa os claims do JWT
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// Server representa o servidor HTTP
type Server struct {
	store      *UserStore
	jwtSecret  []byte
	csrfSecret []byte
	limiter    *IPRateLimiter
}

// NewServer cria um novo servidor
func NewServer() *Server {
	// Gerar segredos aleatórios
	jwtSecret := make([]byte, 32)
	csrfSecret := make([]byte, 32)
	rand.Read(jwtSecret)
	rand.Read(csrfSecret)

	return &Server{
		store:      NewUserStore(),
		jwtSecret:  jwtSecret,
		csrfSecret: csrfSecret,
		limiter:    NewIPRateLimiter(rate.Every(time.Second), 10),
	}
}

// createToken cria um novo JWT
func (s *Server) createToken(userID, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

// validateToken valida um JWT
func (s *Server) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v",
					token.Header["alg"])
			}
			return s.jwtSecret, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}

// generateCSRFToken gera um token CSRF
func (s *Server) generateCSRFToken() string {
	token := make([]byte, 32)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}

// validateCSRFToken valida um token CSRF
func (s *Server) validateCSRFToken(token, cookieToken string) bool {
	return token == cookieToken
}

// authMiddleware é um middleware de autenticação
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		claims, err := s.validateToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// csrfMiddleware é um middleware CSRF
func (s *Server) csrfMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			token := s.generateCSRFToken()
			http.SetCookie(w, &http.Cookie{
				Name:     "csrf_token",
				Value:    token,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteStrictMode,
			})
		} else {
			cookie, err := r.Cookie("csrf_token")
			if err != nil {
				http.Error(w, "CSRF cookie not found", http.StatusBadRequest)
				return
			}

			token := r.Header.Get("X-CSRF-Token")
			if !s.validateCSRFToken(token, cookie.Value) {
				http.Error(w, "Invalid CSRF token", http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

// securityHeadersMiddleware adiciona headers de segurança
func securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; script-src 'self'")
		w.Header().Set("Strict-Transport-Security",
			"max-age=31536000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}

// sanitizeInput sanitiza uma string de entrada
func sanitizeInput(input string) string {
	// Permitir apenas letras, números e alguns caracteres especiais
	reg := regexp.MustCompile(`[^a-zA-Z0-9@._-]`)
	return reg.ReplaceAllString(input, "")
}

// handleLogin processa o login
func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Sanitizar entrada
	username := sanitizeInput(creds.Username)
	password := creds.Password // Não sanitizar senha, pois pode conter caracteres especiais

	// Buscar usuário
	user, ok := s.store.GetByUsername(username)
	if !ok || user.Password != password { // Na prática, usar bcrypt para comparar senhas
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Gerar token
	token, err := s.createToken(user.ID, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

// handleCreateUser cria um novo usuário
func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Sanitizar entrada
	user.Username = sanitizeInput(user.Username)
	// Não sanitizar senha

	// Validar campos
	if len(user.Username) < 3 || len(user.Password) < 8 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verificar se usuário já existe
	if _, exists := s.store.GetByUsername(user.Username); exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Criar usuário
	user.ID = fmt.Sprintf("user_%d", time.Now().UnixNano())
	user.Role = "user" // Role padrão
	s.store.Add(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Criar servidor
	server := NewServer()

	// Criar mux e adicionar rotas
	mux := http.NewServeMux()

	// Rotas públicas
	mux.HandleFunc("/login", server.handleLogin)
	mux.HandleFunc("/users", server.handleCreateUser)

	// Rotas protegidas
	protected := server.authMiddleware(
		server.csrfMiddleware(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Exemplo de rota protegida
				claims := r.Context().Value("claims").(*Claims)
				json.NewEncoder(w).Encode(map[string]string{
					"message": fmt.Sprintf("Hello, %s!", claims.UserID),
				})
			}),
		),
	)
	mux.Handle("/protected", protected)

	// Aplicar middlewares globais
	handler := securityHeadersMiddleware(server.limiter.Middleware(mux))

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         ":8443",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	// Canal para sinais de término
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor em goroutine
	go func() {
		log.Printf("Servidor iniciado em https://localhost%s", srv.Addr)
		if err := srv.ListenAndServeTLS("cert.pem", "key.pem"); err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguardar sinal de término
	<-stop
	log.Println("Desligando servidor...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Erro ao desligar servidor: %v", err)
	}

	log.Println("Servidor desligado")
} 