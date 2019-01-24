package main

import "fmt"

func main() {
	root := NewNode(nil,nil)
	root.SetData("root node")

	a:= NewNode(nil,nil)
	a.SetData("left node")
	b:= NewNode(nil,nil)
	b.SetData("right node")

	root.le =a
	root.ri = b
	fmt.Println(root)
}

type Node struct {
	le *Node
	data interface{}
	ri *Node
}

func NewNode(left,right *Node)*Node{
	return &Node{left,nil,right}
}

func (n *Node)SetData(data interface{}){
	n.data = data
}

