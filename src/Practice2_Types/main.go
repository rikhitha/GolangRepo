package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("types")
	fmt.Println("bool,int8,int16,int32,int64,int")
	fmt.Println("uint8,uint16,uint32,uint64,uint")
	fmt.Println("float32,float64")
	fmt.Println("complex64,complex128")
	fmt.Println("byte")
	fmt.Println("rune,stirng")

	a:=true
	b:=false
	fmt.Println("value is ",a && b,a || b)

	f:=23455678
	g:="hello"
	fmt.Printf("%T %d",f,unsafe.Sizeof(f))
	fmt.Printf("%T,%d",g,unsafe.Sizeof(g))

	c1 := complex(5, 7)
    
    fmt.Println("product:", c1)

	fmt.Println("constants")

	const tg = 45
	fmt.Println(tg)
	

}
