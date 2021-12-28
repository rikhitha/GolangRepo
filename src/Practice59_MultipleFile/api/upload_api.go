package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFiles(response http.ResponseWriter,request *http.Request){

	request.ParseMultipartForm(10*1024*1024)

	files:=request.MultipartForm.File["myfiles"]

	for _,file := range files{

		fmt.Println("File Name" , file.Filename)
		fmt.Println("File Size" , file.Size)
		fmt.Println("File Type" , file.Header.Get("Content-Type"))
		fmt.Println("********************************")

		//Saving files to server

		f,_:=file.Open()

		tempFile,err:=ioutil.TempFile("uploads","upload-*.png")

		if err != nil {
			fmt.Println(err)
		}

		defer tempFile.Close()

		fileBytes,err2:=ioutil.ReadAll(f)

		if err2 !=nil {
			fmt.Println(err)

		}

		tempFile.Write(fileBytes)


	}

	fmt.Println("check uploads folder to view ")
}