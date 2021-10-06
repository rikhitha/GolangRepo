package main

import (
	"fmt"
)

type car struct{
	car string
	enginetype string
}
type carspec struct{
	powerWindows bool
	mileage int
	wheelbase int
	car
}

type carsspec struct{

	carspeci []carspec
}


func (c carsspec) contents(){
	fmt.Println("updations in specification \n ")

	for _,v:=range c.carspeci{
		v.specifications()
		fmt.Println()
	}
}
func (name car)  carname() string{
	return fmt.Sprintf("Car Name := %s ",name.car)
}

func (det carspec) specifications() {
   fmt.Printf(" CarName := %s \n EngineType := %s \n Powerwindows := %t \n Mileage := %d \n WheelBase := %d ",det.carname(),det.enginetype,det.powerWindows,det.mileage,det.wheelbase)
}

func main() {

	fmt.Println("Composition")

	car1:=car{
		"kwidClimber 1.0 AMT",
		"petrol",


	}

	carspec1:=carspec{
		 true ,
		 22,
	    2422,
		car1,


	}

	carspec2:=carspec{
		false ,
		25,
	   2722,
	   car1,


   }
   carspec3:=carspec{
	true ,
	28,
   2722,
   car1,


}
carspec4:=carspec{
 true ,
	22,
   2982,
   car1,


}

	cc:=carsspec{
		carspeci: []carspec{
			carspec1,carspec2,carspec3,carspec4,
		},
	}

	cc.contents()
}

