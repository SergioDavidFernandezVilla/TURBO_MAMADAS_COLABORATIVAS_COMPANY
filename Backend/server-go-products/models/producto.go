package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductoJSON struct {
	ID        string `json:"id"` // es string, porque así llega desde el cliente
	Nombre    string `json:"nombre"`
	Precio    float64 `json:"precio"`
	Stock     int `json:"stock"`
	Categoria string `json:"categoria"`
}

type Producto struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	Nombre             string             `bson:"nombre"`
	Precio             float64             `bson:"precio"`
	Stock              int             `bson:"stock"`
	Categoria          string             `bson:"categoria"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}

type ResponseMessage struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Producto `json:"data"`
}
