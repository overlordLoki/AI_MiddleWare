// src/MemoryGraph/api.go
package memorygraph

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var graph *Graph

// AddNodeHandler adds a new node to the graph.
func AddNodeHandler(c *gin.Context) {
	var requestBody struct {
		Subject string `json:"subject"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node := graph.AddNodeWithUniqueKey(requestBody.Subject)
	c.JSON(http.StatusOK, gin.H{"message": "Node added successfully", "node": node})
}

// GetNodeHandler retrieves a node from the graph.
func GetNodeHandler(c *gin.Context) {
	key := c.Param("key")
	node := graph.FindNode(key)
	if node == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"node": node})
}

// GetGraphHandler retrieves the entire graph.
func GetGraphHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"graph": graph.GetNodes()})
}