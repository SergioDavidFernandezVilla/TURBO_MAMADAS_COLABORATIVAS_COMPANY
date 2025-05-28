package main

import (
	"log"
	"os"

	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-products/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors" // ¡NUEVA IMPORTACIÓN! Necesitas este paquete para CORS
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

	// --- INICIO DE LA SOLUCIÓN CORS ---
	// Configuración de CORS
	config := cors.DefaultConfig()
	// Permite solicitudes desde el origen de tu frontend React (Vite)
	// Asegúrate de que esta URL coincida exactamente con la URL de tu frontend
	config.AllowOrigins = []string{"http://localhost:5173"}

	// Métodos HTTP permitidos (GET, POST, PUT, DELETE, OPTIONS, etc.)
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// Encabezados HTTP permitidos (necesarios si tu frontend envía headers personalizados como Authorization)
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// Si tu frontend necesita enviar cookies o credenciales (tokens de autorización)
	// descomenta la siguiente línea:
	// config.AllowCredentials = true

	// Aplica el middleware CORS al router
	router.Use(cors.New(config))
	// --- FIN DE LA SOLUCIÓN CORS ---

	// Registrar rutas (esto debe ir DESPUÉS de aplicar el middleware CORS)
	routers.Router(router, appCtx)

	// Ejecutar servidor
	log.Printf("Servidor Go de productos escuchando en el puerto :8081") // Mensaje más descriptivo
	router.Run(":8081")
}