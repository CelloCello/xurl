package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestPing(c *gin.Context) {
	resp := Response{}
	resp.Message = "pong"
	c.JSON(http.StatusOK, resp)
}
