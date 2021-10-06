package main

import (
	"fmt"
	"sort"
)

func sorted() {
	fmt.Println("Finished sorting")
}

func sorting(val []int) {

	defer sorted()
	fmt.Println("Values before sorting", val)
	fmt.Println("Started to sort in ascending order")
	sort.Ints(val)
	fmt.Println("sorted values:=", val)

}

type Employee interface {
	getDetails()
	getSalary()
}

type Lead struct {
	name       string
	experience int
	salary     int
}

type SeniorProgrammer struct {
	name       string
	experience int
	salary     int
}

func (l Lead) getDetails() {

	fmt.Printf("Lead Details:= \n name = %s experience = %d \n ", l.name, l.experience)
}

func (s SeniorProgrammer) getDetails() {

	fmt.Printf("Senior Programer Details:= \n name = %s experience = %d \n ", s.name, s.experience)
}

func (l Lead) getSalary() {

	fmt.Printf("Lead Salary:= \n Salary= %d \n ", l.salary)
}

func (s SeniorProgrammer) getSalary() {

	fmt.Printf("Senior programmer Salary:= \n Salary= %d  ", s.salary)
}

func getnumber() int {
	var num int
	fmt.Printf(" \n choose Defer with \n 1. functions \n 2. methods \n 3. arguments evaluation \n 4.stack of defers")
	fmt.Scanln(&num)

	return num
}

func main() {

	fmt.Println("DEFER")

	switch n := getnumber(); {
	case n == 1:
		fmt.Println("defer with functions")
		nums := []int{23, 56, 12, 1, 34, 89}
		sorting(nums)

	case n == 2:
		fmt.Println("Deffered methods")
		lead1 := Lead{
			"Hansy",
			9,
			80000,
		}

		sp1 := SeniorProgrammer{

			"vino",
			2,
			40000,
		}

		var interfaceemp Employee

		interfaceemp = lead1

		interfaceemp.getDetails()

		defer interfaceemp.getSalary() //will be printed in the last when main function returns

		interfaceemp = sp1

		interfaceemp.getDetails()

		interfaceemp.getSalary()

	case n == 3:
		fmt.Println("Arguments evaluation")

		nums := []int{23, 56, 12, 1, 34, 89}
		defer sorting(nums)

		nums = []int{34, 56}

		fmt.Println("values in nums array before defered function call ", nums)

	case n == 4:

		fmt.Println("Stack of defers ,Reverse List of  string \n ")

		listString := []string{"vidhya", "nivya", "sandy", "keerthi"}

		fmt.Println("before defer", listString)

		for _, v := range listString {
			defer fmt.Printf("%s  ", v)
		}

	}

}
