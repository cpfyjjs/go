package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Goo Gooo !!!")
		time.Sleep(2*time.Second)
		c <- true

	}()
	// 阻塞直到有数据可取
	<- c
	fmt.Println("main")
}


