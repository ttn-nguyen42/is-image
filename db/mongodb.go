package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	col *mongo.Collection
}

func NewMongoClient() (*MongoClient, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		panic("missing mongodb uri")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic("missing database name")
	}
	colName := os.Getenv("COLLECTION_NAME")
	if colName == "" {
		panic("missing collection name")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to mongodb")
	}
	defer func() {
		client.Disconnect(context.TODO())
	}()

	collection := client.Database(dbName).Collection(colName)
	return &MongoClient{
		col: collection,
	}, nil
}
