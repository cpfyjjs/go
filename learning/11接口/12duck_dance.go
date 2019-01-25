package main

import "fmt"

// 接口与动态类型
// 类似于Python 和Ruby 这样的动态的语言中的动态类型(duck typing);这意味着对象可以提供的方法被处理，而忽略它们的实际类型：
// 它们能做什么比他们是什么重要
func main() {
	b:= new(Bird)
	DuckDance(b)
}


type IDuck interface{
	Quack()
	Walk()
}

func DuckDance(duck IDuck){
	for i:=0;i<3;i++{
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {

}

func (b Bird)Quack(){
	fmt.Println("I am quacking")
}

func(b Bird)Walk(){
	fmt.Println("I am walking")
}