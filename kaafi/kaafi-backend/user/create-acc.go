package user

import (
	"context"
	"encoding/json"
	"kaafi-backend/datebase"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Createacc(w http.ResponseWriter, r *http.Request) {
	var newUser User
	
	DatebaseConfig,err := datebase.InitDatebase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("there is proplem for setting up the datebase(mongodb)"))
		return
	}
	defer DatebaseConfig.Disconnect(context.Background())

	// Set up the database and collection
	Database := DatebaseConfig.Database(datebase.Database)
	collection := Database.Collection(datebase.Collection)

	err = json.NewDecoder(r.Body).Decode(&newUser)
	userExit, err := searchacc(newUser.Phonenumber)
	if userExit != "" {
		w.WriteHeader(http.StatusAlreadyReported)
		userid, err := json.Marshal(userExit)
		w.Write([]byte("the account already exists"))
		w.Write(userid)
		if err != nil {
			w.Write([]byte("we can't able to translate the user id"))
			return
		}
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newUser.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("there is proplem adding the date to datebase"))
		return
	}
	w.Write([]byte("we created your account"))
	userExit = result.InsertedID.(primitive.ObjectID).Hex()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(userExit)
	w.Write(res)

}
