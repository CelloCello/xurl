package main

import (
	"net/http"
	"short_url/apiserver/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// database
	database.Initialize("data.db")

	// routing
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
