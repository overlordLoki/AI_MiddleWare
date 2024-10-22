package main

import (
	"GoGo/src/Direct"
	g "GoGo/src/MemoryGraph"
	"GoGo/src/config"
	"net/http"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Simple hello world handler
func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	// Initialize configuration
	config.InitConfig()
	
	// Initialize MemoryGraph (AI-related setup)
	g.Init()
	
	// Create a new Gin router
	router := gin.Default()
	
	// Enable CORS middleware to allow all origins
	router.Use(cors.AllowAll())
	
	// Serve static files from the "assets" directory
	router.Static("/assets", "./assets")
	
	// Handle favicon requests to avoid 404s in browsers
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./assets/favicon.ico")
	})	
	// Define your routes
	router.GET("/", hello_world)           // Respond with "hello world"
	router.POST("/chat", Direct.Chat)      // Handle chat requests
	router.POST("/prompt", g.NewPrompt)    // Handle AI prompt generation
	router.POST("/oneshot", Direct.SingleChat) // Handle one-shot chat requests
	// Start the server on the port specified in your configuration
	if err := router.Run("localhost:" + config.Config.Port); err != nil {
		panic(err)
	}
}
