package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Task representa uma tarefa
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// TaskStore é um simples armazenamento em memória
type TaskStore struct {
	tasks  []Task
	nextID int
}

// NewTaskStore cria um novo armazenamento de tarefas
func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

// AddTask adiciona uma nova tarefa
func (s *TaskStore) AddTask(title string) Task {
	task := Task{
		ID:        s.nextID,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	s.tasks = append(s.tasks, task)
	s.nextID++
	return task
}

// GetTasks retorna todas as tarefas
func (s *TaskStore) GetTasks() []Task {
	return s.tasks
}

// TaskHandler gerencia as requisições relacionadas a tarefas
type TaskHandler struct {
	store *TaskStore
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(h.store.GetTasks())

	case "POST":
		var input struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		task := h.store.AddTask(input.Title)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Criar store e handler
	store := NewTaskStore()
	taskHandler := &TaskHandler{store: store}

	// Criar mux e registrar rotas
	mux := http.NewServeMux()
	mux.Handle("/tasks", taskHandler)

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