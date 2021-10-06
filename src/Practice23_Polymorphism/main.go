package main

import (
	"fmt"
)

type Employee interface{

	getDetails()
	getSalary()
}

type Lead struct{
	name string
	experience int
	salary int


}

type SeniorProgrammer struct{

	name string
	experience int
	salary int
	
}

func (l Lead) getDetails() {

	fmt.Printf("Lead Details:= \n name = %s experience = %d \n ",l.name,l.experience)
}

func (s SeniorProgrammer) getDetails() {

	fmt.Printf("Senior Programer Details:= \n name = %s experience = %d \n ",s.name,s.experience)
}


func (l Lead) getSalary() {

	fmt.Printf("Lead Salary:= \n Salary= %d \n ",l.salary)
}

func (s SeniorProgrammer) getSalary() {

	fmt.Printf("Senior programmer Salary:= \n Salary= %d  ",s.salary)
}

func main(){

	fmt.Println("Polymorphism \n ")


	lead1:= Lead{
		"Hansy",
        9,
		80000,
	}

	sp1:=SeniorProgrammer{

		"vino",
		2,
		40000,

	}

	var interfaceemp Employee

	interfaceemp = lead1

	interfaceemp.getDetails()

	interfaceemp.getSalary()

	interfaceemp =sp1

	interfaceemp.getDetails()

	interfaceemp.getSalary()

	}


	


