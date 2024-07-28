package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

const (
    LLM_API_URL = "http://localhost:11434/api/chat"
    MODEL_NAME  = "llama3"
)

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type LLMRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
    Stream   bool      `json:"stream"`
}

func sendChat(messages []Message) (*http.Response, error) {
    data := LLMRequest{
        Model:    MODEL_NAME,
        Messages: messages,
        Stream:   true,
    }
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal request: %w", err)
    }

    resp, err := http.Post(LLM_API_URL, "application/json", bytes.NewReader(jsonData))
    if err != nil {
        return nil, fmt.Errorf("failed to send request to LLM API: %w", err)
    }

    return resp, nil
}

func readChat(resp *http.Response) (string, error) {
    defer resp.Body.Close()

    var output string
    decoder := json.NewDecoder(resp.Body)
    for {
        var body map[string]interface{}
        if err := decoder.Decode(&body); err != nil {
            if err == io.EOF {
                break
            }
            return "", fmt.Errorf("failed to decode response: %w", err)
        }
        if done, ok := body["done"].(bool); ok && done {
            break
        }
        if message, ok := body["message"].(map[string]interface{}); ok {
            if content, ok := message["content"].(string); ok {
                output += content
            }
        }
    }
    return output, nil
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var messages []Message
    if err := json.NewDecoder(r.Body).Decode(&messages); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    resp, err := sendChat(messages)
    if err != nil {
        http.Error(w, "Failed to send chat", http.StatusInternalServerError)
        log.Printf("failed to send chat: %v", err)
        return
    }

    aiResponse, err := readChat(resp)
    if err != nil {
        http.Error(w, "Failed to read chat response", http.StatusInternalServerError)
        log.Printf("failed to read chat response: %v", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"response": aiResponse})
}

func main() {
    http.HandleFunc("/api/chat", chatHandler)
    log.Println("Starting server on :9090")
    log.Fatal(http.ListenAndServe(":9090", nil))
}
