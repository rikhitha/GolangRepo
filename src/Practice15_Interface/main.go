package main

import (
	"fmt"
)

type SalaryCalculator interface {
	calculatesal() int
}
type factorial interface {
	calculate() int
}

type Permanent struct {
	empdid int
	bp     int
	pf     int
	gp     int
}

type contract struct {
	empid int
	bp    int
}

type num int

func (pm Permanent) calculatesal() int {
	return pm.bp + pm.gp + pm.pf
}

func (ct contract) calculatesal() int {
	return ct.bp
}

func (n num) calculate() int {

	sum := 1

	for val := 1; val <= int(n); val++ {
		sum = sum * val
	}
	return sum

}

func getnumber() int {
	var num int
	fmt.Printf(" \n enter number")
	fmt.Scanln(&num)

	return num
}

func ctc(s []SalaryCalculator) {
	total := 0

	for _, v := range s {
		total = total + v.calculatesal()
	}
	fmt.Printf("\n cost to company is : %d ", total)
}

func assertion(a interface{}) {
	v, ok := a.(int)
	fmt.Println("value,status :=", v, ok)
}

func findtype(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("\n %s ::: im string", i)

	case int:
		fmt.Printf("\n  %d im integer", i)

	default:
		fmt.Printf("\n unknown type")
	}
}

func main() {
	fmt.Printf("Interfaces\n")

	fmt.Printf("choose  \n 1.factorial \n 2.CTC	\n 3.TypeAssertion \n 4.Type switch")
	switch n := getnumber(); {
	case n == 1:

		fmt.Printf("\nFactorial")
		var f factorial

		f = num(4)

		fmt.Println(" \n factorial of 4 :", f.calculate())

	case n == 2:
		fmt.Printf("\n cost to company")

		pm1 := Permanent{
			empdid: 1,
			bp:     23000,
			pf:     1000,
			gp:     500,
		}
		pm2 := Permanent{
			empdid: 2,
			bp:     24000,
			pf:     1100,
			gp:     600,
		}
		ct1 := contract{
			empid: 3,
			bp:    10000,
		}
		ct2 := contract{
			empid: 4,
			bp:    20000,
		}

		emp := []SalaryCalculator{pm1, pm2, ct1, ct2}
		ctc(emp)

	case n == 3:
		fmt.Printf("\n type assertion")

		var Interassertion interface{} = 34
		assertion(Interassertion)
		var Interassertion1 interface{} = "hello hai"
		assertion(Interassertion1)

	case n == 4:

		fmt.Printf("\n Type switch")

		findtype("hello hai")

	}

}
