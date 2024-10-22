package subjecttree

import (
	t "GoGo/src/types" // Import the types package
	"github.com/google/uuid"
)

// Node represents a node in a general tree
type ChatNode struct {
	Key      string
	Title    *string
	Summory string // a brief description of the the content of the node
	Children []*ChatNode
	// ConversationHistory represents the conversation history
	ConversationHistory []t.Message
}

// NewNode creates and returns a new node with the given key, optional title, and value.
func NewNode(title string) *ChatNode {
	key := GenerateUniqueKey()
	return &ChatNode{
		Key:     key,
		Title:   &title,
		Children: []*ChatNode{},
		ConversationHistory: []t.Message{},
	}
}

// AddChild adds a new child to the node.
func (n *ChatNode) AddChild(child *ChatNode) {
	n.Children = append(n.Children, child)
}

// generateUniqueKey generates a unique key for a new node
func GenerateUniqueKey() string {
	// Implement a unique key generation logic here
	// This can be as simple as generating a UUID
	return uuid.New().String()
}

// addMessage adds a new message to the node's conversation history.
func (n *ChatNode) AddMessage(role string,promot string) {
	user_promot := t.Message{
		Role:    role,
		Content: promot,
	}
	n.ConversationHistory = append(n.ConversationHistory, user_promot)
}