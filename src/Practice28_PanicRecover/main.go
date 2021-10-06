package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
func slicePanic() {
	n := []int{5, 7, 4}
	fmt.Println(n[4])
	fmt.Println("normally returned from a")
}
func DeferPanic() {
	fmt.Println("Inside DEferPanic method")
	n := []int{5, 7, 4, 6, 7, 8, 9}
	fmt.Println("in 4th position", n[4])
	fmt.Println("normally returned from a")
}
func getnumber() int {
	var num int
	fmt.Printf(" \n choose \n 1:panic // DeferPanic \n 2: Recover \n 3: Error using struct")
	fmt.Scanln(&num)

	return num
}

func main() {

	switch n := getnumber(); {

	case n == 1:

		defer DeferPanic()
		slicePanic()
		fmt.Println("normally returned from main")
		firstName := "Elon"
		fullName(&firstName, nil)
		fmt.Println("returned normally from main")

	case n == 2:

		

	}
}
