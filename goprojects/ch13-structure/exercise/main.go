package main

import (
	"fmt"
	ex1 "goprojects/ch13-exercise/ex1"
	ex3 "goprojects/ch13-exercise/ex3"
	"unsafe"
)

func main() {
	myProduct := ex1.Product{
		Name:         "mac",
		Price:        1000000,
		ReviewSource: 100.00,
	}
	ex1.PrintStruct(myProduct)

	padding1 := ex3.PaddingBefore{}
	padding2 := ex3.PaddingAfter{}

	fmt.Printf("size of padding1: %v\n", unsafe.Sizeof(padding1))
	fmt.Printf("size of padding2: %v\n", unsafe.Sizeof(padding2))
}
