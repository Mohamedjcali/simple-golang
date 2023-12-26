package main

import (
    "fmt"
	"kaafi-backend/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}
func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()
	
	// Define routes and handlers
	router.HandleFunc("/sendotp", user.Sendotp).Methods("GET")
	router.HandleFunc("/createacc", user.Createacc).Methods("POST")
	router.HandleFunc("/getuser/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/checkotp", user.Checkotp).Methods("GET")
	router.HandleFunc("/", HomeHandler).Methods("GET")
	

	// Specify the port to listen on
	

	// Start the server
	log.Printf("Server listening on http://localhost:8080", )
	log.Fatal(http.ListenAndServe(":8080", router))
}