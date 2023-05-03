package main

import (
	add "goprojects/ex23.4/add"
	mul "goprojects/ex23.4/multiple"
)

func main() {
	mul.ReadEq("123 3")
	mul.ReadEq("123 abc")

	add.ReadEq("33 44")
	add.ReadEq("11 ab")
}
