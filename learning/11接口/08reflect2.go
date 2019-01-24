package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 =3.4
	v:= reflect.ValueOf(x)

	fmt.Println("settability of v: ",v.CanSet())

	v = reflect.ValueOf(&x)
	fmt.Println("setability of &v: ",v.CanSet())
	fmt.Println("Type of &v is : ",v.Type())

	v = v.Elem()
	fmt.Println("Element is ",v)
	fmt.Println("settability of v",v.CanSet())

	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
