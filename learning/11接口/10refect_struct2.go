package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23,"skido"}
	s := reflect.ValueOf(&t).Elem()

	//typeofT := s.Type()

	for i:=0;i<s.NumField();i++{
		f:= s.Field(i)
		fmt.Println(f.CanSet())
	}
}
