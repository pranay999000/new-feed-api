package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	feedmodels "github.com/pranay999000/social-minor/services/feeds/feedModels"
	feedresponses "github.com/pranay999000/social-minor/services/feeds/feedResponses"
)

func CreateFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFeed feedmodels.Feed
		c.Bind(&newFeed)
		user_id := c.GetFloat64("id")
		newFeed.UserId = int64(user_id)

		if newFeed.Title != "" {
			feed := newFeed.CreateFeed()

			c.JSON(
				http.StatusCreated,
				feedresponses.FeedResponse {
					Status: http.StatusCreated,
					Message: "success",
					Data: map[string]interface{}{"feed": feed},
				},
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H {
					"message": "title cannot be empty",
				},
			)
		}
	}
}

func GetAllFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		feeds := feedmodels.GetAllFeed()

		c.JSON(
			http.StatusOK,
			feedresponses.FeedResponse {
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"feeds": feeds},
			},
		)
	}
}