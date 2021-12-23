package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

type ContentImage struct {
	ImageName  string `json:"ImageName"`
	Basestring string `json:"BaseString"`
}

var ContentImages []ContentImage

func main() {

	bytes, err := ioutil.ReadFile("./moon.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	// Print the full base64 representation of the image
	fmt.Println(base64Encoding)

	fmt.Println("Rest API v2.0 - Mux Routers")

	ContentImages = []ContentImage{
		{ImageName: "img1.png", Basestring: "hierhgirb"},
	}

	//  str:="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 "

	// databytearray := []byte(str)
	// fmt.Println(databytearray)

	//img,_,err:=image.Decode(bytes.NewReader(databytearray))

	//ioutil.WriteFile("./converted.jpg", data, 0644)

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

