package main

import (
	"net/http" // HTTP package
	"github.com/gin-gonic/gin" // Gin web framework
	cors "github.com/rs/cors/wrapper/gin" // CORS middleware
	"GoGo/src/Direct" // Import the Direct package
	"GoGo/src/config" // Import the config package
)

func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
    config.InitConfig()
    router := gin.Default()
    
    // Configure CORS
    router.Use(cors.AllowAll())
    
    // Serve static files from the "assets" directory
    router.Static("/assets", "./assets")
    
    // Handle the favicon.ico request
    router.GET("/favicon.ico", func(c *gin.Context) {
        c.File("./assets/favicon.ico")
    })
    
    // Other routes
    router.GET("/", hello_world)
    router.POST("/chat", direct.Chat)
    
    router.Run("localhost:" + config.Config.Port)
}
