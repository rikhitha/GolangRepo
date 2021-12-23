package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Createmany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// var datas []user

	datas := []interface{}{}

	err := json.NewDecoder(r.Body).Decode(&datas)

	if err != nil {
		fmt.Println(err)
	}

	insertResults, err := userCollection.InsertMany(context.TODO(), datas)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted all the documents", insertResults)
	json.NewEncoder(w).Encode(insertResults.InsertedIDs)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
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
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"` //value that has to be matched
		City string `json:"city"` // value that has to be modified
	}
	var body updateBody
	e := json.NewDecoder(r.Body).Decode(&body)
	if e != nil {

		fmt.Print(e)
	}
	filter := bson.D{{"name", body.Name}} // converting value to BSON type
	after := options.After                // for returning updated document
	returnOpt := options.FindOneAndUpdateOptions{

		ReturnDocument: &after,
	}
	update := bson.D{{"$set", bson.D{{"city", body.City}}}}
	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOpt)

	var result primitive.M
	_ = updateResult.Decode(&result)

	json.NewEncoder(w).Encode(result)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Println(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func DeleteManyProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"] //get Parameter value as string

	_id, err := primitive.ObjectIDFromHex(params) // convert params to mongodb Hex ID
	if err != nil {
		fmt.Println(err.Error())
	}
	opts := options.Delete().SetCollation(&options.Collation{}) // to specify language-specific rules for string comparison, such as rules for lettercase
	res, err := userCollection.DeleteMany(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	json.NewEncoder(w).Encode(res.DeletedCount) // return number of documents deleted

}

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", CreateProfile).Methods("POST")
	s.HandleFunc("/createmany", Createmany).Methods("POST")
	s.HandleFunc("/getAllUsers", GetAllUsers).Methods("GET")
	// s.HandleFunc("/getUserProfile", GetUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", UpdateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", s)) // Run Server
}
