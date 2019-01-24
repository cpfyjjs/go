package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

func (s Square)Area()float32{
	return s.side*s.side
}

type Circle struct {
	radius float32
}

func(r Circle)Area()float32{
	return r.radius*r.radius*math.Pi
}

type Shaper interface {
	Area()float32
}

func main() {
	var areaIntf Shaper

	sq1 := new(Square)
	sq1.side = 5


	areaIntf = sq1

	if t,ok := areaIntf.(*Square);ok{
		fmt.Printf("The type of areaIntf is %T\n",t)
	}

	if u,ok := areaIntf.(*Circle);ok{
		fmt.Printf("The type of areaIntf is %T\n",u)
	}else{
		fmt.Printf("areaIntf does not contain a variable of type Circle\n")
	}

	//  type switch
	fmt.Println("type-switch")
	switch t:= areaIntf.(type) {
	case *Square:
		fmt.Println(t)
	case *Circle:
		fmt.Println(t)
	case nil:
		fmt.Println("无效")
	default:
		fmt.Println("Un")

	}



	//测试一个值是否实现某个接口

	//sq := Square{4}
	//if _,ok := sq.(Shaper);ok{
	//	fmt.Println("实现了Shaper")
	//}
}
