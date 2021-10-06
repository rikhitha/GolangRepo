package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

type Response struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Location string `json:"Time"`
}

type Person struct {
	Name string `json:"customerName"`

	Age int `json:"customerAge,omitempty"`

	CreditCardnum int `json:"-"`
}

func getnumber() int {

	var num int
	fmt.Printf(" \n choose \n 1: Marshal/Unmarshal \n 2: json:omit \n 3: unmarshal from file")
	fmt.Scanln(&num)

	return num
}

func main() {
	switch n := getnumber(); {

	case n == 1:

		fmt.Println("Json and go ")
		m := Message{"Alice", "Hello", 1294706395881547000}

		b, err := json.Marshal(m)

		if err != nil {

			panic("Data cannot be encoded")

		}

		fmt.Println("encoded data : --", string(b))

		empJsonData := `{"Name":"George Smith","Age":30,"Location":"Newyork, USA"}`

		empBytes := []byte(empJsonData)

		var emp Response

		fmt.Println("Unmarshling")

		json.Unmarshal(empBytes, &emp)

		fmt.Println("name", emp.Name)

		fmt.Println("age", emp.Age)

		fmt.Println("location", emp.Location)

	case n == 2:

		fmt.Println(" \n Marshal")
		p := Person{
			Name: "Tom",
			// Age :22,
			CreditCardnum: 6748394637,
		}

		pbytes, error := json.Marshal(p)

		if error != nil {
			panic("cannot be encoded")

		}

		fmt.Println("values", string(pbytes))

		fmt.Println(" \n UnMarshal")

		pRawJson := json.RawMessage(`{"customerName":"Rikhitha","age":22,"CreditCardnum":5678345}`)

		var pstruct Person

		error2 := json.Unmarshal(pRawJson, &pstruct)

		if error2 != nil {

			panic("cannot decode")

		}

		fmt.Println("unmarshal", pstruct)
		log.Printf(" %+v", pstruct)

	case n == 3:
		fmt.Println("unmarshal from a file")

		fmt.Println(" \n Marshal")
		p := Person{
			Name: "Tom",
			// Age :22,
			CreditCardnum: 6748394637,
		}

		pbytes, error := json.Marshal(p)

		if error != nil {
			panic("cannot be encoded")

		}

		fmt.Println("values", string(pbytes))

		fmt.Println(" \n UnMarshal")

		// pRawJson := json.RawMessage(`{"customerName":"Rikhitha","age":22,"CreditCardnum":5678345}`)

		pRawJson, error3 := ioutil.ReadFile("C:/Users/user/Desktop/name.json")

		if error3 != nil {
			panic("cannot readfile")
		}
		var pstruct Person

		error2 := json.Unmarshal(pRawJson, &pstruct)

		if error2 != nil {

			panic("cannot decode")

		}

		fmt.Println("unmarshal", pstruct)
		log.Printf(" %+v", pstruct)

	}

}
