package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Imagen struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EntidadTipo    string             `bson:"entidad_tipo" json:"entidad_tipo"` // Ej: "producto", "empleado"
	EntidadID      primitive.ObjectID `bson:"entidad_id" json:"entidad_id"`
	NombreArchivo  string             `bson:"nombre_archivo" json:"nombre_archivo"`
	NombreOriginal string             `bson:"nombre_original" json:"nombre_original"`
	Alt            string             `bson:"alt" json:"alt"`
	Url            string             `bson:"url" json:"url"`
	CreacionAt     time.Time          `bson:"creacion_at" json:"creacion_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type ImagenRequest struct {
	EntidadID    string `json:"producto_id"`
	NombreArchivo string `json:"nombre_archivo"`
	NombreOriginal string `json:"nombre_original"`
	Alt           string `json:"alt"`
	Url           string `json:"url"`
}

type ResponseMessage struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []Imagen `json:"data"`
}
