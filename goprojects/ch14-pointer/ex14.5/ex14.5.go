package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("lee", 32)
	fmt.Println(unsafe.Sizeof(userPointer))
	fmt.Println(reflect.TypeOf(userPointer))
}
