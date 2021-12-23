package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ContentImage struct {
	ImageName      byte`json:"ImageName"`
	BaseString string `json:"BaseString"`
	
}

var ContentImages []ContentImage
func main(){
    fmt.Println("Rest API v2.0 - Mux Routers")
	ContentImages = []ContentImage{
		{ImageName:101,BaseString:"agfegfehweigeg" },
		
	}

    handleRequests()
}

func handleRequests() {

    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/getimage", getimages)
    myRouter.HandleFunc("/sendencoding", sendimage).Methods("POST")

    log.Fatal(http.ListenAndServe(":80", myRouter))
}

func sendimage(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	// reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))

	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	fmt.Println("Endpoint Hit: post")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var contentimage ContentImage
	json.Unmarshal(reqBody, &contentimage)
	// update our global Articles array to include
	// our new Article
	ContentImages = append(ContentImages, contentimage)

	json.NewEncoder(w).Encode(contentimage)
}

func getimages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " \n You 've requested : %s \n !", r.URL.Path)
	fmt.Println("Endpoint Hit: getimages")
	json.NewEncoder(w).Encode(ContentImages)
}