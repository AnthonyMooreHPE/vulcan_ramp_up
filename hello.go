package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// api routing
	router := gin.Default()

	// GET requests
	router.GET("/", TestCall)
	router.GET("/hello", HelloWorld)

	// PORT
	router.Run(":80")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Hello": "World",
	})
}

func TestCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
