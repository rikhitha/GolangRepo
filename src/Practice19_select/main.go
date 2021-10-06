package main

import (
	"fmt"
	"time"
)

func func1(ch chan string){
	time.Sleep(2*time.Second)
	ch <-"function1"
}
func func2(ch chan string){
	time.Sleep(11*time.Second)
	ch <-"function2"
}
func func3(ch chan string){
	time.Sleep(1*time.Second)
	ch <-"function3"
}
func main() {

	fmt.Println("Select")

	out1:=make(chan string)

	out2:=make(chan string)

	out3:=make(chan string)

	go func1(out1)

	go func2(out2 )

	go func3(out3 )
	
	select {
	case s1:=<-out1:
		fmt.Println(s1)
	case s2:= <-out2:
		fmt.Println(s2)
	case s3:= <-out3:
		fmt.Println(s3)

	}
}
