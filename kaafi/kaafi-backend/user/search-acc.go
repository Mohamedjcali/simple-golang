package user

import (
	"context"
	"kaafi-backend/datebase"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func searchacc(phoneNumber string) (string, error) {
	DatebaseConfig, err := datebase.InitDatebase()
	if err != nil {
		return "" ,err
	}
	defer DatebaseConfig.Disconnect(context.Background())

	// Set up the database and collection
	Database := DatebaseConfig.Database(datebase.Database)
	collection := Database.Collection(datebase.Collection)

	// Query for the user based on phone number
	filter := bson.M{"phonenumber": phoneNumber}
	var user User
	err = collection.FindOne(context.Background(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		// User not found
		return "", nil
	} else if err != nil {
		// Other error
		return "", err
	}

	// User found, return the user's ID
	return user.Id.Hex(), nil
}
