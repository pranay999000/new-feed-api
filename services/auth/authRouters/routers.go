package authrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/services/auth/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/auth/login", controllers.LoginUser())
	r.POST("api/auth/signup", controllers.CreateUser())
}