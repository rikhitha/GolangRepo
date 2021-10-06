package main

import (
	"fmt"
	"math"
)

func simpleIntrest(amount, intrest, time float64) (string, float64) {
	process := "Simple intrest"
	sim := (amount * intrest * time) / 100

	return process, float64(sim)

}
func compoundIntrest(pamt, ir, duration float64) (cif, ci float64, proc string) {

	proc = "compound intrest"
	cif = pamt * (math.Pow((1 + ir/100), duration))

	ci = cif - pamt

	return

}

func main() {
	process, si := simpleIntrest(1100000, 9.5, 7)
	fmt.Println(process, " is ", si)

	cif, ci, _ := compoundIntrest(700000, 9.5, 7)

	fmt.Println("Compound Intrest", ci)
	fmt.Println("future Compound Intrest", cif)
}
