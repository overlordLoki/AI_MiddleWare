package Direct

import (
	"bytes"
	"encoding/json"
	"net/http"

	"GoGo/src/config"
	"GoGo/src/types"
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
