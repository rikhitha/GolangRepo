package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerRepositoryDb struct {
	client *mongo.Client
}

var userCollection = NewCustomerRepositoryDb().Database("goTest").Collection("users")

func (d CustomerRepositoryDb) Findall(w http.ResponseWriter, r *http.Request) ([]Customer, error) {

	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)

	return nil, nil

}

func NewCustomerRepositoryDb() *mongo.Client {
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

	// return client

	return client
}
