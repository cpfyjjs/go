package main

import "fmt"

type TwoInts struct{
	a int
	b int
}

func main() {
	two := &TwoInts{1,2}
	fmt.Print(two.AddThem())
	fmt.Println()
	fmt.Print(two.AddToParam(3))
}

func (this *TwoInts)AddThem() int {
	return this.a + this.b
}

func (this *TwoInts)AddToParam(param int) int {
	return this.a+this.b + param
}