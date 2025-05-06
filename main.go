package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sknavilehal/go-url-shortener/handler"
	"github.com/sknavilehal/go-url-shortener/store"
)

func main() {
	fmt.Printf("Hello Go URL Shortener !🚀")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go url shortner!",
		})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Add specific route for favicon.ico
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No content
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := router.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
