package api

import (
	"fmt"
	"net/http"
	"short_url/apiserver/pkg/database"
	"short_url/apiserver/pkg/shortener"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateLink(c *gin.Context) {
	db := database.GetDB()
	link := database.Link{}

	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "no binding data",
			"error": err.Error(),
		})
		return
	}

	link.ID = uuid.New()
	link.Code = shortener.GenerateCode(link.ID)

	r := db.Create(&link)
	if r.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "failed to create link",
		})
		return
	}

	fmt.Printf("get link: %s\n", link.Url)
	fmt.Printf("create code: %s, id: %s", link.Code, link.ID.String())

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func GetLinks(c *gin.Context) {
	db := database.GetDB()
	links := []database.Link{}

	result := db.Find(&links)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "failed to get links",
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "success",
		"payload": links,
	})
}
