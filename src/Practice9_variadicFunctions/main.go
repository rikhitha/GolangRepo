package main

import (
	"fmt"
)
func multiple(val int,values ...int ){

	for _,v :=range values{
		if v%val==0{
			fmt.Printf("%d is  divisible by %d \n",v,val)
			

		}else{
		fmt.Printf("%d is not divisible by %d \n",v,val)
		}
	}

}

func main(){
	multiple(2,3,4,5,6,7,8,10)

}