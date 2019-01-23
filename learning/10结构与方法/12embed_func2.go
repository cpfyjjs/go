package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	Log
}

//在类型中嵌入功能
func main() {
	  c := new(Customer)
	  c.Name = "Barak Obama"

	  c.msg = "1 - Yes I can!"
	  c.Add("2 - After me the world be a better place")
	  fmt.Println(c.String())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}


