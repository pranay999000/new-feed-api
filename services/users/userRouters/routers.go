package userrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/services/users/controllers"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/api/users", controllers.GetAllUsers())
}