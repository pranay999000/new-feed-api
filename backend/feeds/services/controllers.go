package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/feeds/models"
)

func GetFeeds() gin.HandlerFunc {
	return func(c *gin.Context) {
		feeds := models.GetFeeds()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"feeds": feeds,
		})
	}
}

func CreateFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFeed models.Feed
		c.Bind(&newFeed)
		user_id := c.GetString("id")
		newFeed.UserId = user_id

		if newFeed.Title != "" && len(newFeed.Body) > 250 && len(newFeed.Title) < 30 {
			feed := newFeed.CreateFeed()
			models.CreateRecent(int64(feed.ID))

			c.JSON(http.StatusCreated, gin.H{
				"success": true,
				"data": feed,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Either title is null or length of title is long ot length of body is short",
			})
		}
	}
}

func LikeFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var like models.Like
		c.Bind(&like)

		err := like.CreateLike()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}

func GetRecents() gin.HandlerFunc {
	return func(c *gin.Context) {
		recents := models.GetRecent()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": recents,
		})
	}
}

func UpdateView() gin.HandlerFunc {
	return func(c *gin.Context) {
		feedId := c.Query("feedId")
		f_id, _ := strconv.Atoi(feedId)
		channel := make(chan models.Feed, 1)
		go models.CheckFeed(int64(f_id), channel)

		feed := <-channel

		if (models.Feed{}) == feed {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "feed not found",
			})
			return
		} else {
			models.CreateView(int64(f_id), feed.ViewCount)
			feed.ViewCount += 1
			models.UpdatePopular(feed)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		}
	}
}

func GetPopular() gin.HandlerFunc {
	return func (c *gin.Context) {
		feeds := models.GetPopular()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": feeds,
		})
	}
}