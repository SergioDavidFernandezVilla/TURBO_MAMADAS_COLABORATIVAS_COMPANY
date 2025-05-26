package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/Backend/utils/models"
	image "github.com/SergioDavidFernandezVilla/TURBO_MAMADAS_COLABORATIVAS_COMPANY/server-go-images/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetImagesByProductID(appCtx *models.AppContext) gin.HandlerFunc {
	return func(router *gin.Context) {
		productID := router.Param("product_id")

		// Obtener env
		var database = os.Getenv("DATABASE")
		var collection = os.Getenv("COLLECTION")

		coleccion := appCtx.MongoClient.Database(database).Collection(collection)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := coleccion.Find(ctx, bson.M{"producto_id": productID})
		if err != nil {
			router.JSON(http.StatusInternalServerError, image.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al obtener las imágenes",
				Data:    nil,
			})
		}

		defer cursor.Close(ctx)

	}
}


func UploadToImageProduct(appCtx *models.AppContext) gin.HandlerFunc{
	return func(router *gin.Context){
			// Datos Form
			productIDStr := router.Param("product_id")
			alt := router.PostForm("alt")

			productID, err := primitive.ObjectIDFromHex(productIDStr)
			if err != nil {
				router.JSON(http.StatusBadRequest, image.ResponseMessage{
					Status:  http.StatusBadRequest,
					Message: "ID de producto inválido",
					Data:    nil,
				})
				return
			}

			// Obtener archivo
			file, header, err := router.Request.FormFile("file")
			if err != nil {
				router.JSON(http.StatusBadRequest, image.ResponseMessage{
					Status:  http.StatusBadRequest,
					Message: "Error al obtener el archivo",
					Data:    nil,
				})
				return
			}
			

			file.Close()

			// Obtener ruta raíz del proyecto
			rootPath := "../../images_products"
			

			// Crear un nombre unico para el archivo
			filename := fmt.Sprintf("%s-%s", primitive.NewObjectID().Hex(), header.Filename)

			// Construir ruta absoluta: /ruta/absoluta/del/proyecto/uploads/productos/archivo.png
			saveDir := filepath.Join(rootPath, "uploads", "productos", productID.Hex())
			savePath := filepath.Join(saveDir, filename)

			// Crear directorio si no existe
			if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
				router.JSON(http.StatusInternalServerError, image.ResponseMessage{
					Status:  http.StatusInternalServerError,
					Message: "Error al crear el directorio de imágenes",
					Data:    nil,
				})
				return
			}


			// Guardar el archivo
			out, err := os.Create(savePath)
			if err != nil {
				router.JSON(http.StatusInternalServerError, image.ResponseMessage{
					Status:  http.StatusInternalServerError,
					Message: "Error al guardar el archivo",
					Data:    nil,
				})
				return
			}
			defer out.Close()

			if _, err = io.Copy(out, file); err != nil {
				router.JSON(http.StatusInternalServerError, image.ResponseMessage{
					Status:  http.StatusInternalServerError,
					Message: "Error al guardar el archivo",
					Data:    nil,
				})
				return
			}

			// Crear documento Imagen
			imagen := image.Imagen{
			ID:            primitive.NewObjectID(),
			ProductoID:    productID,
			NombreArchivo: filename,
			NombreOriginal: header.Filename,
			Alt:           alt,
			Url: "/uploads/productos/" + productID.Hex() + "/" + filename,
			CreacionAt:    time.Now(),
			UpdatedAt:     time.Now(),
		}

		fmt.Println("Nombre original:", header.Filename)
			
		// Insertar en la base de datos
		var database = os.Getenv("DATABASE")
		var collection = os.Getenv("COLLECTION")

		coleccion := appCtx.MongoClient.Database(database).Collection(collection)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err = coleccion.InsertOne(ctx, imagen)

		if err != nil {
			router.JSON(http.StatusInternalServerError, image.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al guardar la imagen",
				Data:    nil,
			})
			return
		}

		router.JSON(http.StatusOK, image.ResponseMessage{
			Status:  http.StatusOK,
			Message: "Imagen guardada correctamente",
			Data:    []image.Imagen{imagen},
		})
	}	
}