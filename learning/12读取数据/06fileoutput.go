package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile,outputError := os.OpenFile("04.txt",os.O_WRONLY|os.O_CREATE,0666)
	if outputError != nil{
		fmt.Println("An error occurred with file opening or creation")
		return
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world\n"

	for i:=0;i<10;i++{
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
