package main

import (
	"fmt"
)

type Student struct{
	name , rno string

	age int 

}
type Person struct {  
    string
    int
}

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main(){

	s1:=Student{
		"Rikhitha","ra1832241020007",23,
	}



	fmt.Printf("name :: %s,,rollnumber :: %s ,, age :: %d",s1.name,s1.rno,s1.age)


	emp3 := struct {
        firstName string
        lastName  string
        age       int
        salary    int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)

	p1 := Person{
        string: "naveen",
        int:    50,
		
    }
    fmt.Println(p1.string)
    fmt.Println(p1.int)

    emp := &Employee{
        firstName: "Sam",
        lastName:  "Anderson",
        age:       55,
        salary:    6000,
    }
    fmt.Println("First Name:", (*emp).firstName)
    fmt.Println("Age:", (*emp).age)
}
