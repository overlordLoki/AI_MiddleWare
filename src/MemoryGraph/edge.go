// src/MemoryGraph/edge.go
package memorygraph

// Edge represents an edge between two nodes in the graph.
type Edge struct {
	From     *SubjectNode
	To       *SubjectNode
	Relation float32 // Relationship type or description
}

// NewEdge creates a new edge between two nodes.
func NewEdge(from, to *SubjectNode, relation float32) *Edge {
	return &Edge{
		From:     from,
		To:       to,
		Relation: relation,
	}
}
