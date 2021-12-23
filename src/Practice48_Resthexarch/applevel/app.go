package applevel

import (
	"Practice48_Resthexarch/domain"
	"Practice48_Resthexarch/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	fmt.Println("Server started")
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	//creating routes  with default multiplexer
	// http.HandleFunc("/greet", hello)
	// http.HandleFunc("/getall", getall)
	// http.HandleFunc("/xmlgetall", xmlgetall)
	// http.HandleFunc("/choosexmljson", choosexmljson)

	// log.Fatal(http.ListenAndServe("localhost:8080", nil))

	//without relying default multiplexer

	// router.HandleFunc("/getall", getall).Methods(http.MethodGet)

	// router.HandleFunc("/greet", hello).Methods(http.MethodGet)
	// router.HandleFunc("/createcustomer", createcustomer).Methods(http.MethodPost)

	// router.HandleFunc("/xmlgetall", xmlgetall).Methods(http.MethodGet)
	router.HandleFunc("/choosexmljson", ch.choosexmljson).Methods(http.MethodGet)
	// router.HandleFunc("/getbyid/{cus_id:[0-9]+}", getcustomer).Methods(http.MethodGet)
	// router.HandleFunc("/getbyids/{cus_id}", getcustomerid).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
