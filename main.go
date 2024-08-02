package main

import (
	"net/http" // HTTP package
	"github.com/gin-gonic/gin" // Gin web framework
	cors "github.com/rs/cors/wrapper/gin" //CORs middleware
	"GoGo/src/Direct" // Import the Direct package
)

func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	router := gin.Default()
	// Configure CORS
	router.Use(cors.AllowAll())
	router.GET("/", hello_world)
	router.POST("/chat", direct.Chat)
	router.Run("localhost:8085")
}
