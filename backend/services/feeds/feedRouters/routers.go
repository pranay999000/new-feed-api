package feedrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/services/feeds/controllers"
)

func FeedRouters(r *gin.Engine) {
	r.POST("/api/feed/create", controllers.CreateFeed())
	r.GET("/api/feeds", controllers.GetAllFeed())
	r.POST("/api/feed/upload/image", controllers.UploadImage())
	r.GET("/api/feed/user", controllers.GetFeedByUser())
}