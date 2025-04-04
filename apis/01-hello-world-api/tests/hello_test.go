package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hello-world/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TestHelloHandler(t *testing.T) {
	// Configurar o modo de teste do Gin
	gin.SetMode(gin.TestMode)

	// Criar um roteador de teste
	r := gin.Default()
	r.GET("/hello", handlers.HelloHandler)

	// Criar uma requisição de teste
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	// Executar a requisição
	r.ServeHTTP(w, req)

	// Verificar o código de status
	if w.Code != http.StatusOK {
		t.Errorf("Esperado status code %d, recebido %d", http.StatusOK, w.Code)
	}

	// Verificar o corpo da resposta
	var response map[string]string
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Não foi possível decodificar a resposta: %v", err)
	}

	expectedMessage := "Hello, World!"
	if message, exists := response["message"]; !exists || message != expectedMessage {
		t.Errorf("Esperada mensagem '%s', recebida '%s'", expectedMessage, message)
	}
} 