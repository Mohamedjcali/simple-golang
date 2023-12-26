package datebase

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Database   = "user"
	Collection = "users"
	Url = "mongodb://localhost:27017"
	
)


func InitDatebase() (*mongo.Client, error) {
	var Client *mongo.Client
	Client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(Url))
	if err != nil {
		return nil, err
	}
	return Client, nil
}
