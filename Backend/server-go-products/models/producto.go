package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Producto struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre             string             `bson:"nombre" json:"nombre"`
	Precio             string             `bson:"precio" json:"precio"`
	Stock              string             `bson:"stock" json:"stock"`
	Categoria          string             `bson:"categoria" json:"categoria"`
	FechaCreacion      time.Time          `bson:"fechaCreacion" json:"fecha_Creacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion" json:"fecha_Actualizacion"`
}

type ResponseMessage struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Producto `json:"data"`
}
