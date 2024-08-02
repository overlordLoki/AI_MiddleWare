package direct

import (
	"GoGo/src/types"
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

const LLM_API_URL = "http://localhost:11434/api/chat"
const MODEL_NAME = "llama3.1"
func sendChatToLLM(messages []types.Message) (*http.Response, error) {
	data := types.ChatRequest{
		Model:    MODEL_NAME,
		Messages: messages,
		Stream:   true,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(LLM_API_URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func readLLMResponse(response *http.Response) (string, error) {
	defer response.Body.Close()

	var output string
	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		var body map[string]interface{}
		err = json.Unmarshal(line, &body)
		if err != nil {
			return "", err
		}
		if done, ok := body["done"].(bool); ok && !done {
			if message, ok := body["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					output += content
				}
			}
		} else {
			break
		}
	}

	return output, nil
}

func Chat(c *gin.Context) {
	var request struct {
		Messages []types.Message `json:"messages"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := sendChatToLLM(request.Messages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	aiResponse, err := readLLMResponse(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.Messages = append(request.Messages, types.Message{Role: "assistant", Content: aiResponse})
	c.JSON(http.StatusOK, types.ChatResponse{Response: aiResponse})
}