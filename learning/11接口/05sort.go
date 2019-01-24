package main

import "fmt"

func main() {
	data := []int{3,1,2,6,7,2,4,5,6,7,0}
	a := IntArray(data)
	Sort(a)
	fmt.Println(a)

	strs :=[]string{"monday","friday","tuesday","wednesday","sunday","thursday","","saturday"}
	b := StringArray(strs)
	Sort(b)
	fmt.Println(b)


}

// 定义一个Sorter接口
type Sorter interface {
	Len() int
	Less(i,j int)bool
	Swap(i,j int)
}

// 定义一个方法，用于为实现该接口的结构排序
func Sort(data Sorter)  {
	for pass :=1;pass<data.Len();pass++{
		for i:=0;i<data.Len()-pass;i++{
			if data.Less(i+1,i){
				data.Swap(i,i+1)
			}
		}
	}
}

func IsSorted(data Sorter)bool{
	n := data.Len()
	for i:=n-1;i>0;i--{
		if data.Less(i,i-1){
			return false
		}
	}
	return true
}


// int 类型的切片实现
type IntArray []int

func (p IntArray)Len()int{
	return len(p)
}

func (p IntArray)Less(i,j int)bool{
	return p[i]<p[j]
}

func (p IntArray)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}


// string类型的切片实现
type StringArray []string

func (p StringArray)Len()int{
	return len(p)
}

func (p StringArray)Less(i,j int)bool{
	return p[i] < p[j]
}

func (p StringArray)Swap(i,j int){
	p[i],p[j] = p[j],p[i]
}

