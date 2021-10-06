package main

import (
	"fmt"
)

func loopone() {
	for j := 10; j >= 0; j = j - 2 {
		fmt.Printf("%d ", j)
	}
}
func breakst() {
	fmt.Printf("\nBREAK")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("\n%d ", i)
	}
	fmt.Printf("\nline after for loop")
}
func continuest() {
	fmt.Printf("\nCONTINUE")
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Printf("\n%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

func nested() {
	fmt.Printf("\nNested ..Pattern\n")

	for g := 0; g < 5; g++ {
		for p := 0; p <= g; p++ {

			fmt.Print("$")
		}
		fmt.Println()
	}

}
func label() {
	fmt.Printf("\nouter label\n")

outer:
	for h := 10; h > 0; h-- {
		for d := 11; d > 0; d-- {
			fmt.Printf("h=%d d=%d", h, d)
			if h == d {
				break outer

			}
		}
	}

}
func forvariation() {
	fmt.Printf("\n for variation\n")

	no, i := 10, 1

	for i <= 10 {
		fmt.Printf("%d X %d = %d \n ",no,i ,no*i)
		i++
	}

}
func main() {
	loopone()
	breakst()
	continuest()
	nested()
	label()
	forvariation()

}
