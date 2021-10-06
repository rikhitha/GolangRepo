package switchh

import (
	"fmt"
)

func getnumber() int {
	var num int
	fmt.Println("enter number")
	fmt.Scanln(&num)

	return num
}

func newSwitch() {

	switch num := getnumber(); { //num is not a constant
	case num < 300:
		fmt.Printf("%d is lesser than 300\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

}
