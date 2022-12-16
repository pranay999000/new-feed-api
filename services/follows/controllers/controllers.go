package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	followmodels "github.com/pranay999000/social-minor/services/follows/followModels"
	followresponses "github.com/pranay999000/social-minor/services/follows/followResponses"
)

func CreateFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFollow followmodels.Follow
		c.Bind(&newFollow)
		user_id := c.GetFloat64("id")
		newFollow.FollowerId = int64(user_id)

		if !followmodels.CheckFollow(int64(user_id), newFollow.FollowingId) {
			f := newFollow.CreateFollow()

			c.JSON(
				http.StatusCreated,
				followresponses.FollowResponse {
					Status: http.StatusCreated,
					Message: "success",
					Data: map[string]interface{}{"follow": f},
				},
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H {
					"message": "already following this user",
				},
			)
		}
	}
}

func GetFollowers() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetFloat64("id")
		user_id := int64(id)
		myId := c.DefaultQuery("user_id", string(rune(user_id)))
		uId, _ := strconv.Atoi(myId)

		followers := followmodels.GetFollowers(int64(uId))

		c.JSON(
			http.StatusOK,
			followresponses.FollowResponse {
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"followers": followers},
			},
		)

	}
}

func GetFollowings() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetFloat64("id")
		user_id := int64(id)
		myId := c.DefaultQuery("user_id", string(rune(user_id)))
		uId, _ := strconv.Atoi(myId)

		followings := followmodels.GetFollowing(int64(uId))

		c.JSON(
			http.StatusOK,
			followresponses.FollowResponse {
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"followings": followings},
			},
		)
	}
}

func DeleteFollow() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetFloat64("id")
		following_id := c.DefaultQuery("user_id", "0")
		fId, _ := strconv.Atoi(following_id)

		if int64(fId) == 0 {
			c.JSON(
				http.StatusBadRequest,
				gin.H {
					"message": "invalid user_id",
				},
			)
		} else {
			f := followmodels.DeleteFollow(int64(id), int64(fId))

			if f {
				c.JSON(
					http.StatusOK,
					gin.H {
						"message": "successfully deleted follow",
					},
				)
			} else {
				c.JSON(
					http.StatusOK,
					gin.H {
						"message": "unable to delete follow",
					},
				)
			}
		}

	}
}