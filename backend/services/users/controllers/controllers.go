package controllers

import (
	"net/http"
	"strconv"

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

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.DefaultQuery("user_id", "0")
		user_id, _ := strconv.Atoi(id)
		user, db := usermodel.GetUserById(int64(user_id))
		var requested usermodel.User
		c.Bind(&requested)

		if db.RecordNotFound() {
			c.JSON(http.StatusNotFound, responses.UserResponse{
				Status: http.StatusNotFound,
				Message: "user not found",
			})
		} else {
			if requested.Name != "" {
				user.Name = requested.Name
			}

			if requested.Image != "" {
				user.Image = requested.Image
			}

			db := usermodel.UpdateUser(user)

			if db.RowsAffected == 1 {
				c.JSON(http.StatusAccepted, responses.UserResponse{
					Status: http.StatusAccepted,
					Message: "success",
					Data: map[string]interface{}{"user": user},
				})
			} else {
				c.JSON(http.StatusAccepted, responses.UserResponse{
					Status: http.StatusBadRequest,
					Message: "failure",
					Data: map[string]interface{}{"user": nil},
				})
			}

		}
	}
}