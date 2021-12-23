package applevel

import (
	"Practice49_DbHexarch/service"
	"encoding/json"
	"encoding/xml"

	"net/http"
)

type Customer struct {
	Customerid string `json:"id" xml:"id"`
	Name       string `json:"YourName" xml:"name"`
	City       string `json:"YourCity" xml:"city"`
	Zipcode    string `json:"YourZipcode" xml:"zipcode"`
}

var Cus []Customer

type CustomerHandlers struct {
	service service.CustomerService
}

// func getall(w http.ResponseWriter, r *http.Request) {

// 	customers := []Customer{
// 		{Customerid: "A2930", Name: "Ashish", City: "erode", Zipcode: "600100"},
// 		{Customerid: "A2931", Name: "Abi", City: "Chennai", Zipcode: "600100"},
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(customers)
// }
// func hello(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "helloo")

// }
// func xmlgetall(w http.ResponseWriter, r *http.Request) {
// 	cus := []Customer{
// 		{Name: "xmlAshish", City: "Chennai", Zipcode: "600100"},
// 		{Name: "xmlAbi", City: "Chennai", Zipcode: "600100"},
// 	}
// 	// w.Header().Add("Content-Type", "application/json")
// 	w.Header().Add("Content-Type", "application/xml")
// 	// json.NewEncoder(w).Encode(customers)
// 	xml.NewEncoder(w).Encode(cus)
// }

func (ch *CustomerHandlers) choosexmljson(w http.ResponseWriter, r *http.Request) {

	// cuschoose := []Customer{
	// 	{Name: "vinay", City: "kanchipuram", Zipcode: "600123"},
	// 	{Name: "preethi", City: "Chennai", Zipcode: "600100"},
	// }

	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")

		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}

// func getcustomerid(w http.ResponseWriter, r *http.Request) {

// 	customers := []Customer{
// 		{Customerid: "A2930", Name: "Ashish", City: "erode", Zipcode: "600100"},
// 		{Customerid: "A2931", Name: "Abi", City: "Chennai", Zipcode: "600100"},
// 	}

// 	vars := mux.Vars(r)

// 	fmt.Fprint(w, vars["cus_id"])

// 	value := vars["cus_id"]

// 	for _, v := range customers {

// 		if v.Customerid == value {
// 			json.NewEncoder(w).Encode(v)
// 		}

// 	}
// 	fmt.Println(w, "Endpoint Hit: returned customer of id:: "+value)

// }
// func getcustomer(w http.ResponseWriter, r *http.Request) {

// 	// customers := []Customer{
// 	// 	{Customerid: "A2930", Name: "Ashish", City: "erode", Zipcode: "600100"},
// 	// 	{Customerid: "A2931", Name: "Abi", City: "Chennai", Zipcode: "600100"},
// 	// }

// 	vars := mux.Vars(r)

// 	fmt.Fprint(w, vars["cus_id"])

// 	// value := vars["cus_id"]

// 	// for _, v := range customers {

// 	// 	if v.Customerid == value {
// 	// 		json.NewEncoder(w).Encode(v)
// 	// 	}

// 	// }
// 	// fmt.Println(w, "Endpoint Hit: returned customer of id:: "+value)

// }

// func createcustomer(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "post request")

// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var customer Customer

// 	json.Unmarshal(reqBody, &customer)

// 	Cus = append(Cus, customer)

// 	json.NewEncoder(w).Encode(Cus)
// 	fmt.Println("st", Cus)
// }
