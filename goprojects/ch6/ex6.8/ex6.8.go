package main

import (
	"fmt"
	"math"
)

func equal(x, y float64) bool {
	return math.Nextafter(x, y) == y
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("a = %0.18f\n", a)
	fmt.Printf("b = %0.18f\n", b)
	fmt.Printf("c = %0.18f\n", c)

	fmt.Printf("%0.18f + %0.18f == %0.18f : %v\n", a, b, c, equal(c, a+b))
}
