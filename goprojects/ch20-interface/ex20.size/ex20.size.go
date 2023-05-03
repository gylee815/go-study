package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Reader interface {
	Read()
}

type ValFile struct {
	Name string
}

type PointerFile struct {
	Name string
}

func (f ValFile) Read() {
	fmt.Println("File Read")
}

func (f *PointerFile) Read() {
	fmt.Println("File Read")
}

func main() {
	var reader1 Reader
	fmt.Printf("reader1 => Type: %T, Size: %v\n\n", reader1, unsafe.Sizeof(reader1))

	reader1 = ValFile{"GYLEE"}
	fmt.Println("reader1 is implemented by ValFile")
	fmt.Printf("reader1 => Type: %T, Size: %v\n\n", reader1, unsafe.Sizeof(reader1))

	var reader2 Reader
	fmt.Printf("reader2 => Type: %T, Size: %v\n\n", reader2, unsafe.Sizeof(reader2))

	pFile := &PointerFile{"GYLEE"}
	reader2 = pFile
	fmt.Println("reader2 is implemented by PointerFile")
	fmt.Printf("reader2 => Type: %v, Size: %v\n\n", reflect.TypeOf(reader2), unsafe.Sizeof(reader2))

	fmt.Printf("Size of pFile: %v\n", unsafe.Sizeof(pFile))
}
