package main

import "fmt"

func main() {
	var value int
	var isPresent bool
	map1 := make(map[string]int)
	map1["New Delhi"]=55
	map1["Beijing"]= 20
	map1["Washington"]=25
	value,isPresent = map1["Beijing"]
	if isPresent{
		fmt.Printf("The value of \"Beijing\" in map1 is %d  \n",value)
	}else{
		fmt.Println(" map1 does not contain Beijing")
	}

	value ,isPresent = map1["Paris"]
	fmt.Println(isPresent)
	fmt.Println(value)

	delete(map1,"Washington")

}
