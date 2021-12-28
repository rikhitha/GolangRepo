package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

// A Response struct to map the Entire Response
// type Response struct {
//     Name    string    `json:"name"`
//     Pokemon []Pokemon `json:"pokemon_entries"`
// }

// // A Pokemon Struct to map every pokemon to.
// type Pokemon struct {
//     EntryNo int            `json:"entry_number"`
//     Species PokemonSpecies `json:"pokemon_species"`
// }

// // A struct to map our Pokemon's Species which includes it's name
// type PokemonSpecies struct {
//     Name string `json:"name"`
// }

type Response struct {
	ImageName      byte`json:"ImageName"`
	BaseString string `json:"BaseString"`
	
}

var ContentImages []Response

func main() {
    response, err := http.Get("http://localhost:80/getimage")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject)

	

	store:=string(responseData)

	fmt.Printf(store)

    fmt.Printf("%v",string(responseData))
    // fmt.Println(len(responseObject.ImageName))

    // for i := 0; i < len(responseObject.Pokemon); i++ {
    //     fmt.Println(responseObject.Pokemon[i].Species.Name)
    // }

}