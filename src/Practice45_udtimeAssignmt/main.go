package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Time struct {
	Current_Time string `json:"CurrentTime"`
}

func main() {

	fmt.Println("Time api")

	router := mux.NewRouter()

	router.HandleFunc("/api/time", gettime)

	fmt.Println("Starting server on port 5000....")
	if error := http.ListenAndServe("localhost:5000", router); error != nil {
		log.Fatalf("not able to start the server : %s ", error.Error())
	}

}

func gettime(w http.ResponseWriter, r *http.Request) {

}
