package main

import "fmt"

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition)getValue()float32{
	return s.sharePrice*s.count
}

type car struct{
	make string
	model string
	price float32
}

func (c car)getValue()float32{
	return c.price
}

type valueable interface {
	getValue()float32
}

func showValue(asset valueable){
	fmt.Printf("Value og the asset is %f\n",asset.getValue())
}


func main() {
	var v valueable = stockPosition{"Glood",577.0,20}
	showValue(v)

	o := car{"BWM","M3",65000}
	showValue(o)

	
}
