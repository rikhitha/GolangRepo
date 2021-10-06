package main

import (
	"errors"
	"fmt"
	"math"
)

func getnumber() int {
	var num int
	fmt.Printf(" \n choose \n 1:radius \n 2: factorial \n 3: Error using struct")
	fmt.Scanln(&num)

	return num
}
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func fact(num int, factchan chan int) (int, error) {

	total := 0

	for o := 1; o <= num; o++ {

		total = total * o
	}
	if total == 0 {
		return 0, fmt.Errorf("Factorial calculation failed Total value is %d ", total)

	}

	return total, nil

}

type ageError struct {
	err string
	age int
}

func (e *ageError) Error() string {
	return fmt.Sprintf(" age : %d : %s", e.age, e.err)
}

func ageValidation(age int) (int, error) {
	if age < 18 {
		return 0, &ageError{"registration not open for this age", age}
	}
	return age, nil
}

func main() {

	switch n := getnumber(); {
	case n == 1:
		radius := -20.0
		area, err := circleArea(radius)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Printf("Area of circle %0.2f", area)

	case n == 2:

		factchan := make(chan int)

		total, error := fact(4, factchan)

		if error != nil {
			fmt.Println(error)
			return
		}

		fmt.Println(" factorial of 4 ", total)

	case n == 3:

		fmt.Printf("Error rep using struct \n")

		var age int

		fmt.Println("Enter age")
		fmt.Scanln(&age)

		ageIsValid, err := ageValidation(age)

		if err != nil {
			if err, ok := err.(*ageError); ok {
				fmt.Printf("Age : %d is not considered", err.age)
				return
			}
			fmt.Println(err)
			return
		}

		fmt.Printf("Age :%d validated", ageIsValid)

	}
}
