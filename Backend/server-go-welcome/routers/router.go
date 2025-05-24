package routers

import (
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router.GET("/", handlers.Welcome)
}