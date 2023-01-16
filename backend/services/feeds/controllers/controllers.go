package controllers

import (
	"net/http"
	"os/exec"
	"path/filepath"

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

func GetFeedByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.GetFloat64("id")
		feeds := feedmodels.GetFeedByUser(int64(user_id))

		c.JSON(http.StatusOK, feedresponses.FeedResponse {
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"feed": feeds},
		})
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

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		f, err := c.FormFile("image_file")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": err.Error(),
			})
		}

		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"error": err.Error(),
			})
		}

		extension := filepath.Ext(f.Filename)
		
		if err := c.SaveUploadedFile(f, "feedImage/" + string(uuid[:len(uuid) - 2]) + extension); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to upload file!",
				"error": err.Error(),
			})
			return
		}

		c.JSON(
			http.StatusOK,
			feedresponses.FeedResponse {
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"image": "/feedImage/" + string(uuid[:len(uuid) - 2]) + extension},
			},
		)
	}
}