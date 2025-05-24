package main

import (
	"log"
	"os"

	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-products/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error al cargar .env: %v", err)
	}

	// Obtener env
	var mongoURL = os.Getenv("MONGO_URL")

	if mongoURL == "" {
		log.Fatalf("❌ No se encontró la variable de entorno MONGO_URL")
	}

	// Crear MongoClient y AppContext
	client := utils.ConnectMongoDB(mongoURL)
	appCtx := &models.AppContext{
		MongoClient: client,
	}

	// Iniciar router
	router := gin.Default()

	// Registrar rutas
	routers.Router(router, appCtx)

	// Ejecutar servidor
	router.Run(":8081")
}
