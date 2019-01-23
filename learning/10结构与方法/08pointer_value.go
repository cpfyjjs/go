package main

import "fmt"

type B struct{
	thing int
}

func (b *B)change(){
	b.thing = 1
}

func (b B)write(){
	fmt.Println(b.thing)
}

func main() {
	var b B
	b.thing =2
	b.write()
	b.change()
	b.write()

	b2 := &B{3}
	b2.change()
	b2.write()
}
