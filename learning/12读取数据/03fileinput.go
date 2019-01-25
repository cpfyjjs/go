package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile,inputError := os.Open("test.txt")

	if inputError != nil{
		fmt.Println("open error",inputError)
		return
	}

	defer inputFile.Close()

	inputRead := bufio.NewReader(inputFile)
	for {
		inputString,readerError := inputRead.ReadString('\n')
		fmt.Println(inputString)
		if readerError == io.EOF{
			return
		}
	}
}
