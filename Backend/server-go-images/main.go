package main

import (
	"log"
	"os"

	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-images/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors" // Importación necesaria para CORS
)

func main() {

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error al cargar .env: %v", err)
	}

	// Obtener la URL de MongoDB desde las variables de entorno
	var mongoURL = os.Getenv("MONGO_URL")

	// Verificar si la variable de entorno MONGO_URL está definida
	if mongoURL == "" {
		log.Fatalf("❌ No se encontró la variable de entorno MONGO_URL")
	}

	// Crear cliente de MongoDB y AppContext
	client := utils.ConnectMongoDB(mongoURL)
	appCtx := &models.AppContext{
		MongoClient: client,
	}

	// Inicializar el router de Gin
	router := gin.Default()

	// --- Configuración e implementación del middleware CORS ---
	// Crear una configuración CORS por defecto
	config := cors.DefaultConfig()
	// Permitir solicitudes desde el origen de tu frontend React (Vite)
	// Asegúrate de que esta URL coincida exactamente con la URL de tu frontend
	config.AllowOrigins = []string{"http://localhost:5173"}

	// Definir los métodos HTTP permitidos para las solicitudes CORS
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	// Definir los encabezados HTTP permitidos (útil si tu frontend envía headers personalizados como Authorization)
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}

	// Si tu frontend necesita enviar cookies o credenciales (tokens de autorización),
	// descomenta la siguiente línea. Úsala con precaución.
	// config.AllowCredentials = true

	// Aplicar el middleware CORS al router de Gin.
	// Es crucial que esto se aplique ANTES de registrar tus rutas.
	router.Use(cors.New(config))
	// --- Fin de la configuración CORS ---

	// Registrar las rutas de la aplicación (esto debe ir DESPUÉS del middleware CORS)
	routers.Router(router, appCtx)

	// Ejecutar el servidor Gin en el puerto 8082
	// Asegúrate de que este es el puerto correcto para el servidor de imágenes
	log.Printf("Servidor Go de imágenes escuchando en el puerto :8082")
	router.Run(":8082")
}
