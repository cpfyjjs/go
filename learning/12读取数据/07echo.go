package main

import (
	"flag"
	"fmt"
)

func main() {
	newline := flag.Bool("n",false,"print newline")
	newint := flag.Int("i",0,"print")
	flag.Parse()
	fmt.Println(*newline)
	fmt.Println(*newint)
}
