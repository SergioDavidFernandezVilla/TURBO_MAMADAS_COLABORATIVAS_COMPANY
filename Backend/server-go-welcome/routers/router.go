package routers

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine){
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}