package main

import (
	"fmt"
	"time"
)

func main() {
	for i:= 0;i<10;i++{
		go Go()
	}
	time.Sleep(2*time.Second)
	print("main")

}

func Go(){
	fmt.Println("Go Go Go !!!")
	time.Sleep(2*time.Second)
}
