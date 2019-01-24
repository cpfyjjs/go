package main

import "bytes"

type ReadWrite interface {
	Read (b bytes.Buffer)bool
	Write (b bytes.Buffer) bool
}

type Look interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Look
	Close()
}

func main() {
	
}
