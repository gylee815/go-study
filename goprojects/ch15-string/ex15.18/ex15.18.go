package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}

	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

func ToUpper3(str string) []byte {
	// var rst string
	slice := []byte(str)
	for i, c := range slice {
		if c >= 'a' && c <= 'z' {
			slice[i] = ('A' + (c - 'a'))
		} else {
			slice[i] = (c)
		}
	}

	return slice
}

func ToUpper3P(str string) *[]byte {
	// var rst string
	slice := []byte(str)
	for i, c := range slice {
		if c >= 'a' && c <= 'z' {
			slice[i] = ('A' + (c - 'a'))
		} else {
			slice[i] = (c)
		}
	}

	return &slice
}

func ToUpper4(str string) string {
	return strings.ToUpper(str)
}

func main() {
	fmt.Println(ToUpper1("giyean.lee"))
	fmt.Println(ToUpper2("giyean.lee"))
	// fmt.Printf("%s\n", *ToUpper3("giyean.lee"))

	func3 := ToUpper3("giyean.lee")
	func3p := ToUpper3P("giyean.lee")

	fmt.Printf("func3 -> Type: %T, Size: %v\n", func3, unsafe.Sizeof(func3))
	fmt.Printf("func3p -> Type: %T, Size: %v\n", *func3p, unsafe.Sizeof(*func3p))
}
