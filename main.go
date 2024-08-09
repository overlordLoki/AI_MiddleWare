package main

import (
	"net/http" // HTTP package
	"github.com/gin-gonic/gin" // Gin web framework
	cors "github.com/rs/cors/wrapper/gin" // CORS middleware
	"GoGo/src/Direct" // Import the Direct package
	"GoGo/src/config" // Import the config package
    "GoGo/src/DeepMemory" // Import the DeepMemory package
    g "GoGo/src/MemoryGraph" // Import the MemoryGraph package
)

func hello_world(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
    config.InitConfig()
    g.InitGraph()
    router := gin.Default()
    
    // Configure CORS
    router.Use(cors.AllowAll())
    
    // Serve static files from the "assets" directory
    router.Static("/assets", "./assets")
    
    // Handle the favicon.ico request
    router.GET("/favicon.ico", func(c *gin.Context) {
        c.File("./assets/favicon.ico")
    })
    deepmemory.InitTree()
    // Other routes
    router.GET("/", hello_world)
    router.POST("/chat", Direct.Chat)
    router.GET("/tree", deepmemory.GetTree)
    router.POST("/memchat", deepmemory.MemChat)
    router.POST("/addnode", deepmemory.AddNode)
    if err := router.Run("localhost:" + config.Config.Port); err != nil {
		panic(err)
	}
}
