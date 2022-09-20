package api

import (
	"log"
	"net/http"
	"xurl/apiserver/pkg/database"
	"xurl/apiserver/pkg/shortener"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateLink(c *gin.Context) {
	db := database.GetDB()
	link := database.Link{}
	resp := CreateLinkResponse{}

	if err := c.ShouldBindJSON(&link); err != nil {
		resp.Message = "no binding data"
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	link.ID = uuid.New()
	link.Code = shortener.GenerateCode(link.ID)

	r := db.Create(&link)
	if r.Error != nil {
		resp.Message = "failed to create link"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	log.Printf("get link: %s\n", link.Url)
	log.Printf("create code: %s, id: %s", link.Code, link.ID.String())

	resp.Message = "success"
	resp.Payload = link
	c.JSON(http.StatusOK, resp)
}

func GetLinks(c *gin.Context) {
	db := database.GetDB()
	links := []database.Link{}

	result := db.Order("created_at desc").Find(&links)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "failed to get links",
			"error": result.Error.Error(),
		})
		return
	}

	resp := LinksResponse{}
	resp.Message = "success"
	resp.Payload = links
	c.JSON(http.StatusOK, resp)
}

func DeleteLink(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("id")

	link := database.Link{}
	resp := Response{}
	if err := db.Where("id = ?", id).First(&link).Error; err != nil {
		resp.Message = "link not found"
		resp.Error = err.Error()
		c.JSON(http.StatusForbidden, resp)
		return
	}

	r := db.Delete(&link)
	if r.Error != nil {
		resp.Message = "failed to delete"
		resp.Error = r.Error.Error()
		c.JSON(http.StatusForbidden, resp)
		return

	}

	resp.Message = "success"
	c.JSON(http.StatusOK, resp)
}
