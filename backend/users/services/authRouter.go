package services

import "github.com/gin-gonic/gin"

func AuthRouters(r *gin.Engine) {
	r.POST("/signup", SignUpUser())
	r.POST("/login", Login())
	r.GET("/list", GetAllUsers())
}