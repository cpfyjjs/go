package main

import "fmt"

func main() {
	// 有缓存是异步的，无缓存同步阻塞的
	c:= make(chan bool,1)

	go func() {
		fmt.Println("Go Go Go")

		<- c
	}()
	c <- true
}
