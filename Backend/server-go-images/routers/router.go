package routers

import (
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-images/handlers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, appCtx *models.AppContext) {
	api := router.Group("/api/v1")
	{
		api.GET("/upload/:tipo/:id", handlers.GetImagesByProductID(appCtx))
		api.POST("/upload/:tipo/:id", handlers.UploadToImageProduct(appCtx))
		api.Static("/uploads", "../../images_products/uploads")
	}
}
