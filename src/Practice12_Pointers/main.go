package main

import (
	"fmt"
)

func change(val *int) *int{
	*val++
	g:=67
	return &g


}
func main() {

	fmt.Println("Pointers")

	d := 25

	var dadd *int = &d

	fmt.Printf("value of d %d ,address of d is  %v ", d, dadd)

	fmt.Println("dereferncing pointers")

	fmt.Printf("value of d %d , value of *dadd:  %d",d,*dadd)

	f:=change(dadd)

	fmt.Printf("after pointer value of d :%d ,value returned : %d",d,f)



}
