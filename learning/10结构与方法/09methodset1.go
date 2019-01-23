package main

import "fmt"

type list []int

func(l list)Len()int{
	return len(l)
}

func(l *list)Append(val int){
	*l = append(*l,val)
}


func main() {
	var lst list
	lst.Append(1)
	lst.Append(2)
	fmt.Print(lst)
}
