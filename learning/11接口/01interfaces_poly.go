package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}

	shapers := []Shaper{r,q}

	for n,_ := range shapers{
		fmt.Println(shapers[n])
		fmt.Println(shapers[n].Area())
	}

}
