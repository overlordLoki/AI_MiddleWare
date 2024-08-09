// src/MemoryGraph/addNode.go
package memorygraph

import (
	"github.com/google/uuid"
)

// generateUniqueKey generates a unique key for a new node
func GenerateUniqueKey() string {
	// Implement a unique key generation logic here
	// This can be as simple as generating a UUID
	return uuid.New().String()
}

// AddNodeWithUniqueKey adds a new node to the graph with a unique key.
func (g *Graph) AddNodeWithUniqueKey(subject string) *SubjectNode {
	key := GenerateUniqueKey()
	node := NewNode(key, subject)
	g.AddNode(node)
	return node
}
