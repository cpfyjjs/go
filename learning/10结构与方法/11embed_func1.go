package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

//在类型中嵌入功能
func main() {
	  c := new(Customer)
	  c.Name = "Barak Obama"
	  c.log = new(Log)

	  c.log.msg = "1 - Yes I can!"
	  c.Log().Add("2 - After me the world be a better place")
	  fmt.Println(c.log.String())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer)Log()*Log{
	return c.log
}
