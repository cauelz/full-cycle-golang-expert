package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler retorna uma mensagem de Hello World
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
} 