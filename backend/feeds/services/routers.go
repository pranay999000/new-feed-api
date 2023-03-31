package services

import "github.com/gin-gonic/gin"

func FeedRouters(r *gin.Engine) {
	r.GET("/feeds", GetFeeds())
}