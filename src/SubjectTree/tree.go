// src/SubjectTree/tree.go
package subjecttree



type Tree struct {
	Title string
	Root *ChatNode
}

//add a node to the tree
func (t *Tree) AddNode(node *ChatNode) {
	if t.Root == nil {
		t.Root = node
		return
	}
	t.Root.AddChild(node)
}

func NewTree(title string) *Tree {
	return &Tree{
		Title: title,
	}
}