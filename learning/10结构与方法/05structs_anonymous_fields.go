package main

import "fmt"

type innerS struct{
	in1 int
	in2 int
}

type outerS struct{
	b int
	c float32
	int
	innerS
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c =3.14
	outer.int = 8
	outer.in1 = 12
	outer.innerS.in2 = 14

	fmt.Print(outer)
//	&{6 3.14 8 {12 14}}
}
