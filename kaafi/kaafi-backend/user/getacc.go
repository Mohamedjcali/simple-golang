package user

import (
	"context"
	"encoding/json"
	"kaafi-backend/datebase"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User represents the structure of a user document in MongoDB

func GetUser(w http.ResponseWriter, r *http.Request) {
	DatebaseConfig, err := datebase.InitDatebase()
	Database := DatebaseConfig.Database(datebase.Database)
	collection := Database.Collection(datebase.Collection)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("there is proplem connecting the datebase"))
		return
	}
	// Get the "id" parameter from the URL
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Println(userID)
	filter := bson.M{"phonenumber": userID}
	var user User
	err = collection.FindOne(context.Background(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("there is No document which have that id "))
		fmt.Fprint(w ,userID)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("there is eror for searching your id"))
		return
	}

	// Convert user to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
