package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"

	// "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type ContentImage struct {
	ImageName      string `json:"ImageName"`
	// BaseString string `json:"BaseString"`
	
}

var ContentImages []ContentImage

func main(){


	f, err := os.Open("./moon.jpg")
  fmt.Print("file open----",f)
if err != nil {
    fmt.Print(err)
}
defer f.Close()

img, fmtName, err := image.Decode(f)
if err != nil {
  fmt.Print("Error",err)
}
fmt.Print("img",img)

// str1 := string(img[:])

// fmt.Print(str1)

// fmt.Print("img",img)
fmt.Print("fmtName",fmtName)


//Encoding(Saving the image)
fn, err := os.Create("./outimage.jpg")
if err != nil {
    // Handle error
}
defer fn.Close()

// Specify the quality, between 0-100
// Higher is better
opt := jpeg.Options{
    Quality: 90,
}
err = jpeg.Encode(fn, img, &opt)
if err != nil {
    // Handle error
}


    fmt.Println("Rest API v2.0 - Mux Routers")
	ContentImages = []ContentImage{
		{ImageName:"101"},
		// {BaseString:"agfegfehweigeg" },
		
	}

	// Image Encoding
	

//call handle requests

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