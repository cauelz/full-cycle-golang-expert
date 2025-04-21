package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Domain types
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Errors
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// Service interface
type UserService interface {
	List() ([]User, error)
	Get(id string) (User, error)
	Create(user User) error
	Update(id string, user User) error
	Delete(id string) error
}

// SQLite implementation
type SQLiteUserService struct {
	db *sql.DB
}

func NewSQLiteUserService(db *sql.DB) *SQLiteUserService {
	return &SQLiteUserService{db: db}
}

func (s *SQLiteUserService) List() ([]User, error) {
	rows, err := s.db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *SQLiteUserService) Get(id string) (User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return u, errors.New("user not found")
	}
	return u, err
}

func (s *SQLiteUserService) Create(user User) error {
	result, err := s.db.Exec(
		"INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)",
		user.Name, user.Email, time.Now(),
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func (s *SQLiteUserService) Update(id string, user User) error {
	result, err := s.db.Exec(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		user.Name, user.Email, id,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (s *SQLiteUserService) Delete(id string) error {
	result, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

// Handler
type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (h *UserHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		h.respondJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
	case errors.As(err, &ValidationError{}):
		h.respondJSON(w, http.StatusBadRequest, err)
	default:
		log.Printf("Error: %v", err)
		h.respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List()
	if err != nil {
		h.handleError(w, err)
		return
	}
	h.respondJSON(w, http.StatusOK, users)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := h.service.Get(id)
	if err != nil {
		h.handleError(w, err)
		return
	}
	h.respondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.handleError(w, &ValidationError{Field: "body", Message: "invalid JSON"})
		return
	}

	if user.Name == "" {
		h.handleError(w, &ValidationError{Field: "name", Message: "required"})
		return
	}
	if user.Email == "" {
		h.handleError(w, &ValidationError{Field: "email", Message: "required"})
		return
	}

	if err := h.service.Create(user); err != nil {
		h.handleError(w, err)
		return
	}
	h.respondJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.handleError(w, &ValidationError{Field: "body", Message: "invalid JSON"})
		return
	}

	if err := h.service.Update(id, user); err != nil {
		h.handleError(w, err)
		return
	}
	h.respondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.service.Delete(id); err != nil {
		h.handleError(w, err)
		return
	}
	h.respondJSON(w, http.StatusNoContent, nil)
}

func (h *UserHandler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", h.List)
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("GET /users/{id}", h.Get)
	mux.HandleFunc("PUT /users/{id}", h.Update)
	mux.HandleFunc("DELETE /users/{id}", h.Delete)
	return mux
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Configurar banco de dados
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			created_at DATETIME NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Criar serviço e handler
	userService := NewSQLiteUserService(db)
	userHandler := NewUserHandler(userService)

	// Configurar rotas
	mux := http.NewServeMux()
	mux.Handle("/api/", loggingMiddleware(userHandler.Routes()))

	// Configurar servidor
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Canal para sinais do sistema
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor em goroutine
	go func() {
		log.Printf("Servidor iniciado em http://localhost%s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguardar sinal de término
	<-done
	log.Print("Servidor está encerrando...")

	// Criar contexto com timeout para shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Tentar shutdown gracioso
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Servidor forçado a encerrar: %v", err)
	}

	log.Print("Servidor encerrado com sucesso")
} 