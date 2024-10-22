package Direct

import (
	"bytes"
	"encoding/json"
	"net/http"

	"GoGo/src/config"
	"GoGo/src/types"

	"github.com/gin-gonic/gin"
)

// Oneshot handles one-shot generation requests to the LLM API.
func Oneshot(prompt string) (string, error) {
	// Construct the payload for the LLM API request
	data := map[string]interface{}{
		"model":  config.Config.ModelName,
		"prompt": prompt,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Send the request to the LLM API
	resp, err := http.Post(config.Config.LLMAPIURL+"/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response from the LLM API
	var result types.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Return the generated text
	return result.Response, nil
}

// api request for one-shot generation
func SingleChat(c *gin.Context) {
	// Get the prompt from the request
	var request struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a response using the LLM API
	response, err := Oneshot(request.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the generated response
	c.JSON(http.StatusOK, gin.H{"response": response})
}