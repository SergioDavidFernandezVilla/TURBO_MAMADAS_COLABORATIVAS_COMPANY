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
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetImagesByProductID(appCtx *models.AppContext) gin.HandlerFunc {
	return func(router *gin.Context) {
		entidadTipo := router.Param("tipo")
		entidadIDStr := router.Param("id")

		// Convertir a ObjectID
		entidadID, err := primitive.ObjectIDFromHex(entidadIDStr)
		if err != nil {
			router.JSON(http.StatusBadRequest, image.ResponseMessage{
				Status:  http.StatusBadRequest,
				Message: "ID inválido",
				Data:    nil,
			})
			return
		}

		// Obtener env
		var database = os.Getenv("DATABASE")
		var collection = os.Getenv("COLLECTION")

		coleccion := appCtx.MongoClient.Database(database).Collection(collection)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		filter := bson.M{"entidad_tipo": entidadTipo, "entidad_id": entidadID}

		// Consultar imágenes por producto_id
		cursor, err := coleccion.Find(ctx, filter)
		if err != nil {
			router.JSON(http.StatusInternalServerError, image.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al consultar imágenes",
				Data:    nil,
			})
			return
		}
		defer cursor.Close(ctx)

		var imagenes []image.Imagen
		if err := cursor.All(ctx, &imagenes); err != nil {
			router.JSON(http.StatusInternalServerError, image.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al procesar los resultados",
				Data:    nil,
			})
			return
		}

		router.JSON(http.StatusOK, image.ResponseMessage{
			Status:  http.StatusOK,
			Message: "Imágenes encontradas",
			Data:    imagenes,
		})
	}
}


func UploadToImageProduct(appCtx *models.AppContext) gin.HandlerFunc {
	return func(router *gin.Context) {
		entidadTipo := router.Param("tipo")
		entidadIDStr := router.Param("id")
		alt := router.PostForm("alt")

		
		entidadID, err := primitive.ObjectIDFromHex(entidadIDStr)
		if err != nil {
			router.JSON(http.StatusBadRequest, image.ResponseMessage{
				Status:  http.StatusBadRequest,
				Message: "ID de entidad inválido",
				Data:    nil,
			})
			return
		}


		file, header, err := router.Request.FormFile("file")
		if err != nil {
			router.JSON(http.StatusBadRequest, image.ResponseMessage{
				Status:  http.StatusBadRequest,
				Message: "Error al obtener el archivo",
				Data:    nil,
			})
			return
		}
		defer file.Close()

		rootPath := "../../images_products"
		filename := uuid.New().String() + filepath.Ext(header.Filename)
		saveDir := filepath.Join(rootPath, "uploads", entidadTipo, entidadID.Hex())
		savePath := filepath.Join(saveDir, filename)

		fmt.Print("savePath: ", savePath)
		fmt.Print("saveDir")

		if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
			router.JSON(http.StatusInternalServerError, image.ResponseMessage{
				Status:  http.StatusInternalServerError,
				Message: "Error al crear el directorio de imágenes",
				Data:    nil,
			})
			return
		}

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

		now := time.Now()
		imagen := image.Imagen{
			ID:             primitive.NewObjectID(),
			EntidadTipo:    entidadTipo,
			EntidadID:      entidadID,
			NombreArchivo:  filename,
			NombreOriginal: header.Filename,
			Alt:            alt,
			Url:            fmt.Sprintf("/uploads/%s/%s/%s", entidadTipo, entidadID.Hex(), filename),
			CreacionAt:     now,
			UpdatedAt:      now,
		}

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
