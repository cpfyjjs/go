package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type",reflect.TypeOf(x))
	v:= reflect.ValueOf(x)
	fmt.Println("value : ",v)
	fmt.Println("kind : ",v.Kind())
	fmt.Println("type : ",v.Type())
	fmt.Println(v.Interface())

}
