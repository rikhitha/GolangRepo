// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"
// )

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// var Articles []Article

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/articles", returnAllArticles)
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }
// func choose() int {

// 	var n int
// 	fmt.Scanln("choose : 1.net/http", &n)

// 	return n
// }
// func handleRequest() {
// 	// creates a new instance of a mux router
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	// replace http.HandleFunc with myRouter.HandleFunc
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/all", returnAllArticles)
// 	// finally, instead of passing in nil, we want
// 	// to pass in our newly created router as the second
// 	// argument
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func main() {

// 	switch n := choose(); {

// 	case n == 1:
// 		Articles = []Article{
// 			{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 			{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 		}
// 		handleRequests()

// 	case n == 2:

// 		fmt.Println("Rest API v2.0 - Mux Routers")
// 		Articles = []Article{
// 			{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 			{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 		}
// 		handleRequest()

// 	}
// }

package main
