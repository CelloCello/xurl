package net

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Config(g *gin.Engine) *gin.Engine {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// corsConfig.AllowHeaders = []string{"Authorization", "Origin"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	g.Use(cors.New(corsConfig))
	return g
}

func Init(root string) *gin.Engine {
	g := gin.Default()
	Config(g)
	RouteInit(g, root)
	return g
}
