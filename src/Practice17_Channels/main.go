package main

import (
	"fmt"
)

func recieve(chan2 chan bool) {
	fmt.Println(" \n inside recieve function")

	chan2 <- true
}

func sum(num int ,sumchan chan int ){

	total:=0
	for i:=0;i<num;i++{

		total=total+i

	}
	sumchan <-total

}
 func fact(num int,factchan chan int){

	total:=1

	for o:=1;o<num;o++{

		total=total*o
	}

	factchan <-total

 }


func main() {
	fmt.Println("Channels")

	var chan1 chan int

	if chan1 == nil {
		fmt.Println(" \n channel is nil declaring it ", chan1)
		chan1 = make(chan int)
		fmt.Printf("type of chaneel chan1 is %T", chan1)
	}

	chan2 := make(chan bool)

	go recieve(chan2)

	<-chan2

	fmt.Println("inside main function")

	sumchan:=make(chan int)

	factchan := make(chan int)

	go fact(3,sumchan)

	go sum(4,factchan)

	sum:= <-sumchan

	factorial:= <-factchan

	fmt.Println("sum of 3 , factorial of 4 ",sum,factorial)



}
