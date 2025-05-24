package routers

import (
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-products/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, appCtx *models.AppContext) {
	api := router.Group("/api")
	{
		api.GET("/productos", handlers.GetProductos(appCtx))
		api.POST("/productos", handlers.CreateProductos(appCtx))
	}
}
