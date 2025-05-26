package models

import "go.mongodb.org/mongo-driver/mongo"

// Mongo
type AppContext struct {
	MongoClient *mongo.Client
}
