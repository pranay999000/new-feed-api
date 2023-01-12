package apis

import (
	"github.com/gin-gonic/gin"
	authrouters "github.com/pranay999000/social-minor/services/auth/authRouters"
	feedrouters "github.com/pranay999000/social-minor/services/feeds/feedRouters"
	followrouters "github.com/pranay999000/social-minor/services/follows/followRouters"
	userrouters "github.com/pranay999000/social-minor/services/users/userRouters"
)

func UserRouters(r *gin.Engine) {
	userrouters.UserRoutes(r)
}

func AuthRouters(r *gin.Engine) {
	authrouters.AuthRoutes(r)
}

func FollowRouters(r *gin.Engine) {
	followrouters.FollowRouters(r)
}

func FeedRouters(r *gin.Engine) {
	feedrouters.FeedRouters(r)
}