package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"GoGo/src/types"
	"github.com/stretchr/testify/assert"
	cors "github.com/rs/cors/wrapper/gin"
	"GoGo/src/config"
	"GoGo/src/Direct"
)

// SetupRouter sets up the Gin router for testing
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.AllowAll())
	router.POST("/chat", direct.Chat)
	return router
}

func TestChat(t *testing.T) {
	router := SetupRouter()

	// Create a request body with example messages
	messages := []types.Message{
		{Role: "user", Content: "Hello, how are you?"},
	}
	chatRequest := types.ChatRequest{
		Model:    config.Config.ModelName,
		Messages: messages,
		Stream:   true,
	}
	jsonValue, _ := json.Marshal(chatRequest)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/chat", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "chat sent to LLM", response["message"])
}
