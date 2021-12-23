package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"YourName" xml:"name"`
	City    string `json:"YourCity" xml:"city"`
	Zipcode string `json:"YourZipcode" xml:"zipcode"`
}

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "helloo")

	})
	http.HandleFunc("/getall", getall)
	http.HandleFunc("/xmlgetall", xmlgetall)
	http.HandleFunc("/choosexmljson", choosexmljson)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func getall(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "Chennai", Zipcode: "600100"},
		{Name: "Abi", City: "Chennai", Zipcode: "600100"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func xmlgetall(w http.ResponseWriter, r *http.Request) {
	cus := []Customer{
		{Name: "xmlAshish", City: "Chennai", Zipcode: "600100"},
		{Name: "xmlAbi", City: "Chennai", Zipcode: "600100"},
	}
	// w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Type", "application/xml")
	// json.NewEncoder(w).Encode(customers)
	xml.NewEncoder(w).Encode(cus)
}

func choosexmljson(w http.ResponseWriter, r *http.Request) {

	cuschoose := []Customer{
		{Name: "vinay", City: "kanchipuram", Zipcode: "600123"},
		{Name: "preethi", City: "Chennai", Zipcode: "600100"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")

		xml.NewEncoder(w).Encode(cuschoose)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cuschoose)

	}
}
