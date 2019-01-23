package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y float64
}

func (p *Point)Abs()float64{
	return math.Sqrt(p.x*p.x+p.y+p.y)
}


type NamePoint struct {
	name string
	Point
}

func main() {
	// 内嵌将一个已存在类型的字段和方法注入到另外一个类型里：匿名字段上的方法晋升为外层类型的方法
	n := NamePoint{"x1",Point{3,4}}
	fmt.Println(n.Abs())

	m := &NamePoint{"x1",Point{3,4}}
	fmt.Println(m.Abs())
}
