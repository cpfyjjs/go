package main

import "fmt"

// 假设我们想获取一个map类型的切片，我们必须使用两次make()函数，第一次分配切片，第二次分配切片中的每一个map元素
func main() {
	items := make([]map[int]int,5)

	for i := range items{
		items[i] = make(map[int]int,1)
		items[i][1]=2
	}

	fmt.Println(items)

	items2	:=	make([]map[int]int,	5)
	for _,item := range items2{
		// 不会改变items2 内部的数据，item 为值类型
		item= make(map[int]int,1)
		item[1] =1
	}

	fmt.Println(items2)

}
