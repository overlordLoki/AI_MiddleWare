// src/MemoryGraph/node.go
package memorygraph

import (
	t "GoGo/src/SubjectTree"
)

// Node represents a node in the subject graph.
type SubjectNode struct {
	Key       string
	Subject   string
	Edges     []*Edge
	Trees	[]*t.Tree
}

// NewNode creates a new node with the given key, subject, and value.
func NewNode(key, subject string) *SubjectNode {
	return &SubjectNode{
		Key:     key,
		Subject: subject,
		Edges:   []*Edge{},
		Trees: []*t.Tree{},
	}
}

// AddEdge
func (n *SubjectNode) AddEdge(edge *Edge) {
	n.Edges = append(n.Edges, edge)
}

// AddTree returns the tree that was added
func (n *SubjectNode) AddTree(title string) *t.Tree {
	tree := t.NewTree(title)
	n.Trees = append(n.Trees, tree)
	return tree
}
