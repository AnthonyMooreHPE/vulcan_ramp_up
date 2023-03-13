package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// api routing
	router := gin.Default()

	// GET requests
	router.GET("/hello", HelloWorld)

	// PORT
	router.Run("localhost:8080")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "World",
	})
}
