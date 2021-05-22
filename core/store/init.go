package store

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	Connection string
	MongoDB    *mongo.Database
}

var Conn DatabaseConnection

func Init(ctx context.Context) error {
	var clientOps options.ClientOptions
	clientOps.ApplyURI(os.Getenv("CONN"))
	client, err := mongo.Connect(ctx, &clientOps)
	if err != nil {
		log.Println("unable to establish connection", err)
		return nil
	}
	Conn.MongoDB = client.Database(os.Getenv("MONGODB_NAME"))
	return nil
}
