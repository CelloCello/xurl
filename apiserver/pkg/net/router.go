package net

import (
	"short_url/apiserver/api"

	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", api.TestPing)
	r.GET("api/links", api.GetLinks)
	r.POST("api/links", api.CreateLink)
	r.DELETE("api/links/:id", api.DeleteLink)
	return r
}
