package main

import (
	"fmt"
	"reflect"
)

type NotknowType struct{
	s1,s2,s3 string
}

func (n NotknowType)String()string{
	return n.s1+"-"+n.s2+"-"+n.s3
}

var secret interface{} = NotknowType{"Ada","Go","Oberon"}


func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	fmt.Println(value)
	fmt.Println(typ)

	knd := value.Kind()

	fmt.Println(knd)

	for i:=0;i<value.NumField();i++{
		fmt.Printf("Fidld %d : %v\n",i,value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
