package router

import (
	"short_url/apiserver/api"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", api.TestPing)
	r.POST("api/links", api.LinkCreate)
	return r
}
