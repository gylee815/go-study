package exinit

import (
	"fmt"
)

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Printf("f() d: %d\n", d)
	return d
}

func PrintD() {
	fmt.Println("d: ", d)
}
