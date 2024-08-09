package subjecttree

import (
	t "GoGo/src/types" // Import the types package
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
func NewNode(key string, title *string) *ChatNode {
	return &ChatNode{
		Key:      key,
		Title:    title,
		Children: []*ChatNode{},
	}
}

// AddChild adds a new child to the node.
func (n *ChatNode) AddChild(child *ChatNode) {
	n.Children = append(n.Children, child)
}