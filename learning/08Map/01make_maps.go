package main

import "fmt"

func main() {
	var mapList map[string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one":1,"two":2}
	mapCreated:=make(map[string]float32)
	mapAssigned = mapList

	mapCreated["key1"]=1
	mapCreated["key2"]=2
	mapCreated["key3"]=3

	mapAssigned["one"]=11
	fmt.Printf("%#v\n",mapList)
	fmt.Printf("%#v\n",mapAssigned)
	fmt.Printf("%#v\n",mapCreated)

	//map[string]int{"one":11, "two":2}
	//map[string]int{"one":11, "two":2}
	//map[string]float32{"key1":1, "key2":2, "key3":3}

}
