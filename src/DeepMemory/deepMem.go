// src/DeepMemory/deepMem.go
package deepmemory

import (
	"fmt"
	"net/http"
	"strings"
	"GoGo/src/Direct"
	"github.com/gin-gonic/gin"
)

var root *Node
var subjects []string


// MemChat handles the main chat functionality.
func MemChat(c *gin.Context) {
	if root == nil {
		InitTree()
	}
	subjects = GetSubjects()

	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the subject of the prompt
	promptSubject, err := handleSubject(request.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Print the prompt and the subject for debugging
	fmt.Println("Prompt: " + request.Prompt)
	fmt.Println("Subject: " + promptSubject)

	// Add the new subject to the tree if it's not already present
	if !subjectExists(promptSubject, subjects) {
		newNode := NewNode(GenerateUniqueKey(), promptSubject, request.Prompt)
		root.AddChild(newNode, 1.0) // Assuming a default relative value
		c.JSON(http.StatusOK, gin.H{"message": "New subject added to the tree", "node": newNode})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Subject already exists"})
	}
}

// handleSubject processes the prompt and assigns or creates a subject using AI.
func handleSubject(prompt string) (string, error) {
	aiPrompt := "Subjects: " + formatSubjects(subjects) +
		" Please assign the following prompt to a subject, if no subject fits then please name a new one." +
		"\nPrompt: " + prompt

	// Call the oneshot function from the direct package
	response, err := Direct.Oneshot(aiPrompt)
	if err != nil {
		return "", err
	}

	return response, nil
}

// formatSubjects formats the subjects into a readable string.
func formatSubjects(subjects []string) string {
	return "[" + strings.Join(subjects, ", ") + "]"
}

// subjectExists checks if a subject already exists in the current subjects list.
func subjectExists(subject string, subjects []string) bool {
	for _, s := range subjects {
		if s == subject {
			return true
		}
	}
	return false
}



// InitTree initializes the tree with the root node, containing the user's information.
func InitTree() {
	if root == nil {
		baseInfo := "User Name is Loki"
		root = NewNode("0", "Root", baseInfo)
	}
}

// GetRoot returns the root node of the tree.
func GetRoot() *Node {
	return root
}
