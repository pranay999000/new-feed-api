package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/social-minor/utils/apis"
	"github.com/pranay999000/social-minor/utils/config"
	bearertoken "github.com/vence722/gin-middleware-bearer-token"
)

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server v1.0.0.0",
		})
	})

	apis.AuthRouters(r)

	r.Use(bearertoken.Middleware(func (token string, c *gin.Context) bool {
		claims, ok := config.ValidateJWT(token)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return false
		}

		id := claims["id"]
		
		c.Set("email", claims["email"])
		c.Set("name", claims["name"])
		c.Set("id", id)

		return true
	}))

	apis.UserRouters(r)
	apis.FollowRouters(r)
	apis.FeedRouters(r)

	log.Fatal(r.Run(":9010"))
}