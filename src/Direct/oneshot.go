package direct

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"GoGo/src/config"
	"GoGo/src/types"
)

// Oneshot handles one-shot generation requests to the LLM API.
func Oneshot(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt" binding:"required"`
	}

	// Bind JSON request to the struct
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Construct the payload for the LLM API request
	data := map[string]interface{}{
		"model":  config.Config.ModelName,
		"prompt": request.Prompt,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request data"})
		return
	}

	// Send the request to the LLM API
	resp, err := http.Post(config.Config.LLMAPIURL+"/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request to LLM API"})
		return
	}
	defer resp.Body.Close()

	// Read the response from the LLM API
	var result types.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response from LLM API"})
		return
	}

	// Return the result to the client
	c.JSON(http.StatusOK, result)
}
