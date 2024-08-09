// src/DeepMemory/addNode.go
package deepmemory

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// AddNode adds a new node to the tree
func AddNode(c *gin.Context) {
	type AddNodeRequest struct {
		ParentKey string `json:"parent_key"`
		Subject   string `json:"subject"`
		Title     string `json:"title,omitempty"`
		Value     string `json:"value"`
	}
	var request AddNodeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parentNode := FindNode(request.ParentKey)
	if parentNode == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent node not found"})
		return
	}

	newNode := NewNode(GenerateUniqueKey(), request.Subject, request.Value)
	if request.Title != "" {
		newNode.SetTitle(request.Title)
	}
	parentNode.AddChild(newNode, 0.5) // You can adjust the relative value as needed

	c.JSON(http.StatusOK, gin.H{"message": "Node added successfully", "node": ConvertNodeToResponse(newNode)})
}


// NodeResponse is a simplified version of Node for JSON responses
type NodeResponse struct {
	Key     string         `json:"key"`
	Subject string         `json:"subject"`
	Title   *string        `json:"title,omitempty"`
	Value   string         `json:"value"`
	Children []NodeResponse `json:"children"`
}

// ConvertNodeToResponse converts a Node to a NodeResponse
func ConvertNodeToResponse(node *Node) NodeResponse {
	if node == nil {
		return NodeResponse{}
	}

	children := make([]NodeResponse, len(node.Children))
	for i, child := range node.Children {
		children[i] = ConvertNodeToResponse(child)
	}

	return NodeResponse{
		Key:      node.Key,
		Subject:  node.Subject,
		Title:    node.Title,
		Value:    node.Value,
		Children: children,
	}
}