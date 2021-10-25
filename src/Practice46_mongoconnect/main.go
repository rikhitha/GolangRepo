package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

var userCollection = db().Database("goTest").Collection("users")

func CreateProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var person user
	err := json.NewDecoder(r.Body).Decode(&person) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}
	insertResult, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	json.NewEncoder(w).Encode(insertResult.InsertedID) // return the mongodb ID of generated document

}
func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", CreateProfile).Methods("POST")
	// s.HandleFunc("/getAllUsers", GetAllUsers).Methods("GET")
	// s.HandleFunc("/getUserProfile", GetUserProfile).Methods("POST")
	// s.HandleFunc("/updateProfile", UpdateProfile).Methods("PUT")
	// s.HandleFunc("/deleteProfile/{id}", DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", s)) // Run Server
}
