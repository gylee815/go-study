package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type Writer interface {
	Write()
}

type File struct {
	Name string
}

func (f *File) Read() {
	fmt.Println("Read File")
}

func (f *File) Close() {
	fmt.Println("Close File")
}

func ReadFile(r Reader) {
	r.Read()
	if c, ok := r.(Closer); ok {
		c.Close()

		f := c.(*File)
		fmt.Printf("File: %s is closed\n", f.Name)
	} else {
		fmt.Printf("Error occured, c is %T type and value is %v\n", c, c)
	}
}

func main() {
	myFile := &File{"MyFile"}
	ReadFile(myFile)
}
