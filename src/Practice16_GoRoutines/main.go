package main

import (
	"fmt"
	"time"
)

func goroute1() {

	fmt.Println("inside goroutine 1 function")
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("inside goroutine 1 function")
		fmt.Printf("  %d", i)

	}
}

func goroute2() {
	fmt.Println("in second go function")
	for i := 10; i < 15; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println("in second go function")
		fmt.Printf("  %d", i)

	}

}

func main() {

	go goroute1()

	go goroute2()
	time.Sleep(1 * time.Second)
	fmt.Println(" main Goroutines")
}
