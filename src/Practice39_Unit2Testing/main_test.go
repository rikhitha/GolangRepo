package main

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(10, 10)

	if result != 100 {
		t.Errorf("Expected 100 for product of 10*10,but got '%d' ", result)
	}
}

type tablecase struct {
	var1 int
	var2 int
	var3 int
}

func TestTableMultiply(test *testing.T) {

	cases := []tablecase{
		{
			var1: 10,
			var2: 10,
			var3: 100,
		},
		{
			var1: 20,
			var2: 10,
			var3: 200,
		},
		{
			var1: 10,
			var2: 30,
			var3: 300,
		},
		{
			var1: 10,
			var2: 40,
			var3: 400,
		},
		{
			var1: 10,
			var2: 78,
			var3: 780,
		},
	}

	for _,val:=range cases{
		returnval:=Multiply(val.var1,val.var2)
		if val.var3!=returnval{
			test.Errorf("Expected %d ,but got %d",val.var3,returnval)
		}
	}

}
