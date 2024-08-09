// src/DeepMemory/node.go
package deepmemory

import (
	t "GoGo/src/types" // Import the types package
	"sync"

	"github.com/google/uuid"
)

// Node represents a node in a general tree
type Node struct {
	Key      string
	Subject  string
	Title    *string
	Value    string
	Children []*Node
	Edges    []*Edge // Edges connected to this node
	// ConversationHistory represents the conversation history
	ConversationHistory []t.Message
}


// Edge represents an edge between two nodes in the tree
type Edge struct {
	Relative float64 `json:"relative"` // Relative represents how related the nodes are
	From     *Node   `json:"from"`
	To       *Node   `json:"to"`
}

// NodeRepresentation represents a node with its subject, title, and children
type NodeRepresentation struct {
	Subject  string               `json:"subject"`
	Title    *string              `json:"title,omitempty"`
	Children []NodeRepresentation `json:"children"`
}

// ConvertNodeToRepresentation converts a Node to its representation
func ConvertNodeToRepresentation(node *Node) NodeRepresentation {
	if node == nil {
		return NodeRepresentation{}
	}

	children := make([]NodeRepresentation, len(node.Children))
	for i, child := range node.Children {
		children[i] = ConvertNodeToRepresentation(child)
	}

	return NodeRepresentation{
		Subject:  node.Subject,
		Title:    node.Title,
		Children: children,
	}
}


// NewNode creates and returns a new node with the given key, subject, optional title, and value.
func NewNode(key, subject, value string) *Node {
	return &Node{
		Key:     key,
		Subject: subject,
		Value:   value,
		Children: []*Node{},
		Edges:    []*Edge{},
		ConversationHistory: []t.Message{},
	}
}

// AddChild adds a child node to the current node and creates an edge between them.
func (n *Node) AddChild(child *Node, relative float64) {
	n.Children = append(n.Children, child)
	edge := &Edge{
		Relative: relative,
		From:     n,
		To:       child,
	}
	n.Edges = append(n.Edges, edge)
	child.Edges = append(child.Edges, edge)
}

// SetTitle sets the title of the node
func (n *Node) SetTitle(title string) {
	n.Title = &title
}

// AddMessage adds a message to the node's conversation history
func (n *Node) AddMessage(msg t.Message) {
	n.ConversationHistory = append(n.ConversationHistory, msg)
}

// GetMessages returns the node's conversation history
func (n *Node) GetMessages() []t.Message {
	return n.ConversationHistory
}

// GetChildren returns the node's children
func (n *Node) GetChildren() []*Node {
	return n.Children
}


// GetNodesBySubject returns all nodes with the given subject
func GetNodesBySubject(subject string) []*Node {
	var nodes []*Node

	var traverse func(n *Node)
	traverse = func(n *Node) {
		if n == nil {
			return
		}
		if n.Subject == subject {
			nodes = append(nodes, n)
		}
		for _, child := range n.Children {
			traverse(child)
		}
	}

	traverse(root)
	return nodes
}

// FindNode finds a node by its key
func FindNode(key string) *Node {
	var findNode func(node *Node, key string) *Node
	findNode = func(node *Node, key string) *Node {
		if node == nil {
			return nil
		}
		if node.Key == key {
			return node
		}
		for _, child := range node.Children {
			found := findNode(child, key)
			if found != nil {
				return found
			}
		}
		return nil
	}

	return findNode(root, key)
}

// generateUniqueKey generates a unique key for a new node
func GenerateUniqueKey() string {
	// Implement a unique key generation logic here
	// This can be as simple as generating a UUID
	return uuid.New().String()
}

// GetSubjects returns all unique subjects in the tree
func GetSubjects() []string {
	subjects := make(map[string]bool)
	var mu sync.Mutex // For thread-safe access to the map

	var traverse func(n *Node)
	traverse = func(n *Node) {
		if n == nil {
			return
		}
		mu.Lock()
		subjects[n.Subject] = true
		mu.Unlock()
		for _, child := range n.Children {
			traverse(child)
		}
	}

	traverse(root)

	subjectList := make([]string, 0, len(subjects))
	for subject := range subjects {
		subjectList = append(subjectList, subject)
	}

	return subjectList
}