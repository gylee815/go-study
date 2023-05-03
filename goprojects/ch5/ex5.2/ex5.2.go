package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 2252365876239.2345345

	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Print("a: ", a, " b: ", b, " f: ", f, "\n")
	fmt.Printf("a: %d, b: %d, f: %f", a, b, f)
}
