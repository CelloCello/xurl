package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloView(c *gin.Context) {
	c.HTML(http.StatusOK, "hello.tmpl", gin.H{
		"title": "Hello PONG!",
	})
}
