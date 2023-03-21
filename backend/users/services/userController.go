package services

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/users/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))
		defer cancel()

		opts := options.Find().SetSkip(int64((page - 1) * limit)).SetLimit(int64(limit))
		filter := bson.D{}

		cursor, err := userCollection.Find(ctx, filter, opts)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
			})
			return
		}

		defer cursor.Close(ctx)
		
		var userList []models.User
		for cursor.Next(ctx) {
			var user models.User

			if err = cursor.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
				})
				return
			}

			userList = append(userList, user)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"users": userList,
		})

	}
}