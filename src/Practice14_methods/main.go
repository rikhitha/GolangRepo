package main

import (
	"fmt"
)

type EmployeeDetails struct {
	empid   string
	empname string
}

type EmployerDetails struct {
	CorporateName string
	Industry      string
}

func (emp EmployeeDetails) DisplayDetails() {
	fmt.Printf(" Employee Details : \n id : %s  \n name: %s ", emp.empid, emp.empname)
}

func (emp EmployerDetails) DisplayDetails() {
	fmt.Printf("\n \n Employer Details : \n Corporate Name : %s  \n Industry: %s ", emp.CorporateName, emp.Industry)
}

func (emp EmployerDetails) ChangeEmployerDetails(newname string) {

	emp.CorporateName = newname
	fmt.Printf("\n\n Inside changeemplyr method (Value receivers)  \n")
	fmt.Printf("\n  Employer Details : \n Corporate Name : %s  \n Industry: %s ", emp.CorporateName, emp.Industry)
}

func (e *EmployeeDetails) chnageempid(id string) {
	e.empid = id
	fmt.Printf("\n\n Inside changeempid method (Pointer receivers)  \n")
	fmt.Printf("\n  Employee Details : \n  Employee Id : %s  \n Employee Name: %s ", e.empid, e.empname)

}

func main() {

	employee1 := EmployeeDetails{
		empid:   "A2920",
		empname: "Rikhitha",
	}

	employee1.DisplayDetails()

	employer1 := EmployerDetails{
		CorporateName: "SrmTech",
		Industry:      "IT service",
	}

	employer1.DisplayDetails()

	employer1.ChangeEmployerDetails("infotech")

	fmt.Printf("\n \n After changeemplrdetails(),name of corporate : %s  ", employer1.CorporateName)

	employee1.chnageempid("A2945")

	fmt.Printf("\n \n After changeempid(),id : %s  ", employee1.empid)

}
