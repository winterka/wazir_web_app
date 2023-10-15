package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "tested",
	})
}

func main() {
	route := gin.Default()
	route.GET("/", testing)
	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
