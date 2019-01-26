package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("03.txt")
	if err != nil {
		panic(err)
	}
	// 程序运行完毕执行关闭文件的操作
	defer file.Close()

	var col1, col2, col3, col4 []string
	for {
		var v1, v2, v3 ,v4 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3,&v4)
		// 一直扫描直到结束
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
		col4 = append(col4, v4)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
	fmt.Println(col4)

}
