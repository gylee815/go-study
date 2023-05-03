package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := 500
	p := &a
	var c *int

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("value of var which p is pointing: %d\n", *p)

	*p = 100
	fmt.Printf("value of a: %d\n", a)

	fmt.Printf("type of p: %v, type of *p: %v\n", reflect.TypeOf(p), reflect.TypeOf(*p))
	fmt.Printf("value of &*p: %v\n", &*p)

	fmt.Printf("size of p: %v, size of *p: %v\n", unsafe.Sizeof(p), unsafe.Sizeof(*p))
	fmt.Printf("value of c: %v\n", c)
	// fmt.Printf("value of *c: %v\n", *c)
	fmt.Printf("sizeof c: %v\n", unsafe.Sizeof(c))
}
