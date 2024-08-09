// src/MemoryGraph/graph.go
package memorygraph

import (
	"sync"
)

// Graph represents the overall structure of the subject graph.
type Graph struct {
	Nodes map[string]*SubjectNode
	mu    sync.Mutex
}

func InitGraph() {
	graph = NewGraph()
	// Add some initial nodes and edges
	node1 := NewNode("1", "User information")
	graph.AddNode(node1)
}


//----UTILS FUNCTIONS------------------------------------//

// NewGraph creates a new graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*SubjectNode),
	}
}

// AddNode adds a new node to the graph.
func (g *Graph) AddNode(node *SubjectNode) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Nodes[node.Key] = node
}

// FindNode finds a node in the graph by its key.
func (g *Graph) FindNode(key string) *SubjectNode {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.Nodes[key]
}

// fund the node by its subject
func (g *Graph) FindNodeBySubject(subject string) *SubjectNode {
	g.mu.Lock()
	defer g.mu.Unlock()
	for _, node := range g.Nodes {
		if node.Subject == subject {
			return node
		}
	}
	return nil
}

// AddEdge adds a new edge to the graph.
func (g *Graph) AddEdge(edge *Edge) {
	g.mu.Lock()
	defer g.mu.Unlock()
	edge.From.AddEdge(edge)
	edge.To.AddEdge(edge)
}

// GetNodes returns all nodes in the graph.
func (g *Graph) GetNodes() map[string]*SubjectNode {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.Nodes
}

// get all the subjects from the graph. each node has a subject
func (g *Graph) GetSubjects() []string {
	g.mu.Lock()
	defer g.mu.Unlock()
	subjects := make(map[string]bool)
	for _, node := range g.Nodes {
		subjects[node.Subject] = true
	}
	var result []string
	for subject := range subjects {
		result = append(result, subject)
	}
	return result
}