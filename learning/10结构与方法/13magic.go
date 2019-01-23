package main

import "fmt"

type Base struct {

}

func (Base) Magic(){
	fmt.Println("base magic")
}

func (self Base) MoreMagic(){
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (self Voodoo)Magic(){
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.MoreMagic()
	//base magic
	//base magic

	v.Magic()
	//  voodoo magic

}
