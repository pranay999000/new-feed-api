package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/services/users/responses"
	usermodel "github.com/pranay999000/social-minor/services/users/userModel"
	"github.com/pranay999000/social-minor/utils/config"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds usermodel.User
		c.Bind(&creds)

		if creds.Email != "" && creds.Password != "" {

			user, db := usermodel.GetUserByEmail(creds.Email)

			if db.RecordNotFound() {
				c.JSON(
					http.StatusBadRequest,
					gin.H {
						"message": "User not found",
					},
				)
			} else {
				if user.Password == creds.Password {
					token, err := config.GenerateJWT(user.Email, user.Name, user.ID)

					if err != nil {
						c.JSON(
							http.StatusInternalServerError,
							gin.H {
								"message": "Unable to create token",
							},
						)
					} else {
						c.JSON(
							http.StatusAccepted,
							responses.UserResponse {
								Status: http.StatusAccepted,
								Message: "success",
								Data: map[string]interface{}{"user": user, "token": token},
							},
						)
					}
				} else {
					c.JSON(
						http.StatusBadRequest,
						gin.H {
							"message": "Incorrect password",
						},
					)
				}
			}

		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H {
					"message": "Invalid credentials",
				},
			)
		}
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser usermodel.User
		c.Bind(&newUser)

		if newUser.Name != "" && newUser.Password != "" && newUser.Email != "" {

			_, db := usermodel.GetUserByEmail(newUser.Email)

			if db.RecordNotFound() {
				u := newUser.CreateUser()
				token, err := config.GenerateJWT(u.Email, u.Name, u.ID)

				if err != nil {
					c.JSON(
						http.StatusInternalServerError,
						gin.H {
							"message": "unable to create token",
						},
					)
				} else {
					c.JSON(
						http.StatusCreated, responses.UserResponse{
							Status: http.StatusCreated,
							Message: "success",
							Data: map[string]interface{}{"user": u, "token": token},
						},
					)
				}
			} else {
				c.JSON(
					http.StatusBadRequest,
					gin.H {
						"message": "Email already exists",
					},
				)
			}


		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H {
					"message": "Invalid user details",
				},
			)
		}
	}
}