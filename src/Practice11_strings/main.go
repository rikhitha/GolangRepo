package main

import (
	"fmt"
)

func main() {
	fmt.Println("Strings")

	name := "Rikhitha Manojkumar"

	fmt.Println("unicode")

	for i := 0; i < len(name); i++ {

		fmt.Printf("at %d position:%x", i, name[i])
	}
	fmt.Println("\n characetrs")
	for i := 0; i < len(name); i++ {

		fmt.Printf("\n at %d position:%c", i, name[i])
	}

}
