package main

import "fmt"

func valSwitch(a, b int) (int, int) {
	return b, a
}

func main() {
	x := 1
	y := 2

	x, y = valSwitch(x, y)
	fmt.Printf("x: %d, y: %d\n", x, y)
}
