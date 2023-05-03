package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1
	// str3 := "Hello world"

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("Type of &str1: %T stringHeader1: %v\n", stringHeader1.Data, stringHeader1)
	fmt.Printf("Type of &str2: %T stringHeader2: %v\n", stringHeader2.Data, stringHeader2)
	fmt.Println(unsafe.Pointer(&str1), unsafe.Pointer(&str2))

	fmt.Printf("Size of stringHeader1.Data: %d, Size of stringHeader1.Len: %d\n", unsafe.Sizeof(stringHeader1.Data), unsafe.Sizeof(stringHeader1.Len))
	fmt.Printf("Size of stringHeader1: %d\n", unsafe.Sizeof(*stringHeader1))
}
