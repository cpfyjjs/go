package main

import (
	"fmt"
)

func main() {
	//runtime.GOMAXPROCS(1)
	c := make(chan bool)

	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<-c05
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}
