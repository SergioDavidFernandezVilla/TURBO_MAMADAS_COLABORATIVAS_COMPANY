package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	once        sync.Once
)

// ConnectMongoDB establece una conexión única (singleton)
func ConnectMongoDB(url string) *mongo.Client {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
		if err != nil {
			log.Fatalf("❌ Error al conectar con MongoDB: %v", err)
		}

		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("❌ No se pudo hacer ping a MongoDB: %v", err)
		}

		log.Println("✅ Conectado exitosamente a MongoDB Atlas")
		MongoClient = client
	})

	return MongoClient
}

// GetMongoCollection devuelve una colección específica reutilizando la conexión
func GetMongoCollection(_ string, db, col string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("❌ MongoClient no está inicializado. Llama a ConnectMongoDB primero.")
	}
	return MongoClient.Database(db).Collection(col)
}
