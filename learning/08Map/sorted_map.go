package main

import (
	"fmt"
	"sort"
)

var barVal = map[string]int{
	"alpha": 34, "bravo": 56, "charlie": 23,
	"delta": 87, "echo": 56, "foxtrot": 12,
	"golf": 34, "hotel": 16, "indio": 87,
	"juliet": 65, "kili": 34, "lima": 98}

// 如果你想为map排序，需要将key(或者value)拷贝到一个切片，在独对切边进行排序
func main() {
	fmt.Println("unsorted: ")
	for k,v := range barVal{
		fmt.Printf("key: %v, Value: %v\n",k,v)
	}

	keys := make([]string,len(barVal))

	i:=0
	for k := range barVal{
		keys[i]= k
		i++
	}
	sort.Strings(keys)

	fmt.Println()
	fmt.Println("sorted")
	for _,k:= range keys{
		fmt.Printf("key: %v, Value: %v\n",k,barVal[k])
	}
}
