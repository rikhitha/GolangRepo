package main

import (
	"fmt"
)

func Getarr() {
	var arSize int

	fmt.Println("enter size of the array")

	fmt.Scanln(&arSize)

	var arr = make([]int, arSize)

	for u := 0; u < arSize; u++ {

		fmt.Printf("enter the value for %d element:", u)
		fmt.Scanf("%d", &arr[u])

	}
	fmt.Println("values in array are :", arr)

}
func printarray() {
	arrayval := [...]string{"Vincy", "mansi", "vidh", "crysty", "densy"}

	for i, v := range arrayval {

		fmt.Printf("at postion %d is : %s", i, v)
		fmt.Printf("\n")
	}
	for _,v:= range arrayval{
		fmt.Println(v)
	}
}
func Modifyslice(arrmod [7]int){

	fmt.Println("\n inside modifyslice()")
	fmt.Println("recieved array",arrmod)

	var slicearrmod []int =arrmod[2:5]

	fmt.Println("slice of array 2:5",slicearrmod)

	for i:=range slicearrmod{
		slicearrmod[i]=slicearrmod[i]+2
	}

	fmt.Println("modification in slice",slicearrmod)

	fmt.Println("after modification ,array values :",arrmod)

	var slice2 []int=arrmod[:]

	fmt.Println(": in slice",slice2)




}
func slice(){
	arr1:=[...]int {12,13,14,15,23,76,34}
	var slicearr1 []int =arr1[2:5]
	fmt.Println("values in array:",arr1)
	fmt.Println("values of arra from 2 to 4 in slice:",slicearr1)

	fmt.Printf("slice capacity: %d and length :%d",len(slicearr1),cap(slicearr1))


	Modifyslice(arr1)

	
}
func main() {
	var op int

	fmt.Printf("choose option 1.(array input from user)\n2.(printing array,blank identifier,range)\n 3.(slices/)")

	fmt.Scanln(&op)

	switch op {
	case 1:
		Getarr()
		break
	case 2:
		printarray()
	case 3:
		slice()
	}
}
