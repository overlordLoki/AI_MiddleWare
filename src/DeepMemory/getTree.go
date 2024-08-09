package deepmemory

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


// GetTree returns the entire tree structure with subjects and titles as JSON
func GetTree(c *gin.Context) {
	if root == nil {
		InitTree()
	}

	treeRepresentation := ConvertNodeToRepresentation(root)
	c.JSON(http.StatusOK, treeRepresentation)
}