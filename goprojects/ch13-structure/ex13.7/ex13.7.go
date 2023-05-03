package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	A int8
	B int
	C int8
	D int
	E int8
	F int
}

type MyStruct2 struct {
	A int8
	C int8
	E int8
	B int
	D int
	F int
}

func main() {
	var a MyStruct
	var b MyStruct2
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
