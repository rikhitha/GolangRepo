package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	myRouter.HandleFunc("/create", newemployee).Methods("POST")

	myRouter.HandleFunc("/delete/{Eid}", deleteemp).Methods("DELETE")

	myRouter.HandleFunc("/fetchbyid/{Eid}", fetchbyid)

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

func fetchbyid(writer http.ResponseWriter, response *http.Request) {

	fmt.Fprintf(writer, "DETAILS OF THE EMPLOYEE:: \n \n")

	fmt.Fprintf(writer, "\n Hello you have requested :- %s \n \n", response.URL.Path)

	variable := mux.Vars(response)

	key := variable["Eid"]

	fmt.Fprintf(writer, "Value passed : "+key)

	for _, em := range Emp {
		if em.Eid == key {
			json.NewEncoder(writer).Encode(em)
		}
	}

	fmt.Println(writer, "Endpoint Hit: returned value of key "+key)
}

func newemployee(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: emplyee post")

	reqBody, _ := ioutil.ReadAll(r.Body)

	var employ Employee

	json.Unmarshal(reqBody, &employ)

	Emp = append(Emp, employ)

	json.NewEncoder(w).Encode(employ)

}

func deleteemp(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Enpoint hit : Delete")

	vars := mux.Vars(r)

	id := vars["Eid"]

	for i, e := range Emp {
		if e.Eid == id {

			Emp = append(Emp[:i], Emp[i+1:]...)

		}
	}
}
