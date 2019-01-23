package main

import "fmt"

//当为结构定义一个alias类型时，此机构体类型和它的alias类型都有相同的底层类型，他们之间可以相互转换
type number struct{
	f float32
}

type nr number

func main() {
	a := number{5.0}
	b := nr{5.0}

	var c = number(b)
	fmt.Print(a,b,c)
}
