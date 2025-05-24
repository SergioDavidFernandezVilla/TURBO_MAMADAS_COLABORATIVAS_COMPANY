// handlers/productos.go
package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	mongo "github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-products/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProductos(appCtx *mongo.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		coleccion := appCtx.MongoClient.Database("company").Collection("productos")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := coleccion.Find(ctx, bson.M{})
		if err != nil {
			log.Printf("❌ Error al obtener productos: %v", err)
			c.JSON(http.StatusInternalServerError, models.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al obtener productos",
				Data:    nil,
			})
			return
		}
		defer cursor.Close(ctx)

		var productos []models.Producto
		if err := cursor.All(ctx, &productos); err != nil {
			log.Printf("❌ Error al decodificar productos: %v", err)
			c.JSON(http.StatusInternalServerError, models.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al decodificar productos",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, models.ResponseMessage{
			Status:  http.StatusOK,
			Message: "Productos obtenidos exitosamente",
			Data:    productos,
		})
	}
}

func CreateProductos(appCtx *mongo.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var nuevoProducto models.Producto

		if err := c.ShouldBindJSON(&nuevoProducto); err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseMessage{
				Status:  http.StatusBadRequest,
				Message: "Error en el formato del JSON: " + err.Error(),
				Data:    nil,
			})
			return
		}

		nuevoProducto.ID = primitive.NewObjectID()
		now := time.Now()
		nuevoProducto.FechaCreacion = now
		nuevoProducto.FechaActualizacion = now

		coleccion := appCtx.MongoClient.Database("company").Collection("productos")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := coleccion.InsertOne(ctx, nuevoProducto)
		if err != nil {
			log.Printf("❌ Error al insertar producto: %v", err)
			c.JSON(http.StatusInternalServerError, models.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al insertar producto",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, models.ResponseMessage{
			Status:  http.StatusCreated,
			Message: "Producto creado exitosamente",
			Data:    []models.Producto{nuevoProducto},
		})
	}
}
