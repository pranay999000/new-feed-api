package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pranay999000/social-minor/services/users/responses"
	usermodel "github.com/pranay999000/social-minor/services/users/userModel"
)

var NewUser usermodel.User

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := usermodel.GetAllUsers()
		c.JSON(
			http.StatusOK, responses.UserResponse{
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"users": users},
			},
		)
	}
}