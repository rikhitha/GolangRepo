package main

import (
	"fmt"
	"reflect"
)

type EmployeeDetails struct {
	empid   string
	empname string
}

type EmployerDetails struct {
	CorporateName string
	Industry      string
}

type order struct {
	ordId      int
	customerId int
}

func main() {

	fmt.Println("Reflection")

	switch n := getnum(); {
	case n == 1:

		fmt.Println("reflect.Type and reflect.Value")

		employee1 := EmployeeDetails{
			empid:   "A2920",
			empname: "Rikhitha",
		}

		employer1 := EmployerDetails{
			CorporateName: "SrmTech",
			Industry:      "IT service",
		}

		fmt.Println("\n employee status")
		query(employee1)
		fmt.Println("\n employer status")
		query(employer1)

	case n == 2:
		fmt.Println("reflect.kind")

		employee1 := EmployeeDetails{
			empid:   "A2920",
			empname: "Rikhitha",
		}

		fmt.Println("\n employee kind status")
		kindfunc(employee1)

	case n == 3:

		fmt.Println("NumField() && Field()")

		employer1 := EmployerDetails{
			CorporateName: "SrmTech",
			Industry:      "IT service",
		}

		nofields(employer1)

	case n == 4:
		fmt.Println("Int() & String()")
		a := 3
		t := reflect.ValueOf(a).Int()
		fmt.Printf("type:%T,,value:%v", t, t)
		p := "Rikhitha"
		r := reflect.ValueOf(p).String()
		fmt.Printf("type:%T,,value:%v", r, r)
	}

}

func getnum() int {
	fmt.Println("choose : \n 1.type & value \n 2.kind(specific kind of type) \n 3.numfield() & field() \n 4.Int() & String()")

	var n int

	fmt.Scanln(&n)

	return n
}

func query(q interface{}) {

	fmt.Println("valueof anf type of")
	t := reflect.TypeOf(q)
	value := reflect.ValueOf(q)

	fmt.Println("type :==", t)

	fmt.Println("value :==", value)

}

func kindfunc(q interface{}) {

	fmt.Println("typeof and kind")

	t := reflect.TypeOf(q)

	k := t.Kind()
	value := reflect.ValueOf(q)

	fmt.Println("type :==", t)

	fmt.Println("type of type  :==", k)

	fmt.Println("value :==", value)

}

func nofields(q interface{}) {

	fmt.Println("no of fields nd ith field")

	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		numfield := v.NumField()
		fmt.Println("Number of fields", numfield)
		for i := 0; i < numfield; i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
