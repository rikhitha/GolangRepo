package main

import (
	"fmt"
)

func getnum() int {

	var a int
	fmt.Println("Enter choice-- \n 1. Anonymous Functions \n 2.calling anonymous function without assinging to a variable \n 3.passing arguments to anonymous functions \n 4.Higher order functions \n 5. Returning functions from another function \n 6. Closures")
	fmt.Scanln(&a)
	return a
}

func main() {

	fmt.Println("FirstClassfunctions")

	switch n := getnum(); {

	case n == 1:
		fmt.Println("Assinging functions to variables--Anonymous Functions")

		factorial := func() {
			total := 1
			num := 4

			fmt.Printf("%T \n ", total)

			for o := 1; o < num; o++ {

				total = total * o
			}

			fmt.Println("Total :=", total)
		}

		factorial()

		fmt.Printf("%T", factorial)

	case n == 2:
		fmt.Println(" \n calling anonymous function without assinging to a variable ")

		func() {
			fmt.Println("im insidee anonymous function")
		}()

	case n == 3:
		fmt.Println("\n Passing Arguments to anonymous functions")

		func(age int) {
			fmt.Printf("My age : = %d ", age)
		}(22)

	case n == 4:

		fmt.Println("Higher order functions--take one or more func as arguments // return function as its result")

		f := func(a, b int) int {
			return a + b
		}
		simple(f)

	case n==5:
		fmt.Println("Returning functions from another functions")

		pass:=returnfunc()
		fmt.Println(pass(2,3))

	case n==6 :
		fmt.Println("Closures")

		valueout:=9

		func(){
			fmt.Printf("value:=%d",valueout)
		}()
		


	}

}
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func returnfunc() func(a,b int) int{
	f:=func(a,b int) int{
		return a*b
	}

	return f
}