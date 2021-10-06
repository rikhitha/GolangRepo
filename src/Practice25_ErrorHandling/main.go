package main

import (
	"errors"
	"fmt"
	"os"
)

func getnumber() int {
	var num int
	fmt.Printf(" \n choose \n 1:Files error \n 2: error type")
	fmt.Scanln(&num)

	return num
}
func main() {
	switch n := getnumber(); {
	case n == 1:

		fmt.Println("Error Handling in go")

		f, err := os.Open("C:/Users/user/Desktop/link.txt")

		if err != nil {

			fmt.Println("Errors:=", err)
			return
		}

		fmt.Println(f.Name(), "opened successfully")

	case n == 2:
		var erro error

		erro = errors.New("Fake error")

		fmt.Println("Hello", erro)

	}
}
