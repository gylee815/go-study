package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d address = %v\n", i, v, &a[i])
	}

	for i, v := range b {
		fmt.Printf("b[%d] = %d address = %v\n", i, v, &b[i])
	}

	b = a

	for i, v := range b {
		fmt.Printf("b[%d] = %d address = %v\n", i, v, &b[i])
	}
}
