package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/feeds/models"
)

func GetFeeds() gin.HandlerFunc {
	return func(c *gin.Context) {
		feeds := models.GetFeeds()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"feeds": feeds,
		})
	}
}