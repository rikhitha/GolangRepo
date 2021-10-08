package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Ename    string `json:"Employee Name"`
	Eid      string `json:"Employee Id"`
	Econtact int    `json:"Employee Contact Number"`
	Edept    string `json:"Employee Department"`
}

var Emp []Employee

func main() {

	fmt.Println("EMPLOYEE API")

	Emp = []Employee{
		{
			Ename:    "Vinay",
			Eid:      "A2930",
			Econtact: 7871831352,
			Edept:    "ccx"},
		{
			Ename:    "vaishali",
			Eid:      "A2939",
			Econtact: 9871831352,
			Edept:    "pes"},
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/fetchall", fetchall)

	myRouter.HandleFunc("/fetchbyid", fetchbyid)

	log.Fatal(http.ListenAndServe(":5000", myRouter))

}

func homePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Welcome to employee information portal")
	fmt.Fprintf(writer, " \n \n \n  Hello you 've requested : %s \n !", request.URL.Path)
	fmt.Println("Endpoint Hit: homePage")

}

func fetchall(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "DETAILS OF ALL EMPLOYEES \n \n")

	fmt.Fprintf(w, " \n Hello you 've requested : %s \n \n \n !", r.URL.Path)

	fmt.Println("Endpoint Hit: Fetchall")

	json.NewEncoder(w).Encode(Emp)

}

func fetchbyid(writer http.ResponseWriter, response http.Request) {

}
