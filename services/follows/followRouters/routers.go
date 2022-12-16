package followrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/services/follows/controllers"
)

func FollowRouters(r *gin.Engine) {
	r.POST("/api/follow/create", controllers.CreateFollow())
	r.GET("/api/follow/followers", controllers.GetFollowers())
	r.GET("/api/follow/followings", controllers.GetFollowings())
	r.DELETE("/api/follow/delete", controllers.DeleteFollow())
}