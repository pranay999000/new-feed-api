package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/follows/services"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Follows@v1.0.0.0",
	})
}

func InitRouters(r *gin.Engine) {
	r.GET("/", rootHandler)
	services.FollowRouters(r)
}

func main() {
	r := gin.Default()

	InitRouters(r)

	err := r.Run(":8002")
	if err != nil {
		log.Fatal(err)
	}
}