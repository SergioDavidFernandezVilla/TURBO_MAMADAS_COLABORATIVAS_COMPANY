package handlers

import (
	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to TURBO MAMADAS COLABORATIVAS COMPANY",
	})
}