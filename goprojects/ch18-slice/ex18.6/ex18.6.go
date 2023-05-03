package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Printf("slice1: %v, len(slice1): %d, cap(slice1): %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, len(slice2): %d, cap(slice2): %d\n", slice2, len(slice2), cap(slice2))
	fmt.Println()

	slice1[1] = 100
	fmt.Println("After change second element of slice1")
	fmt.Printf("slice1: %v, len(slice1): %d, cap(slice1): %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, len(slice2): %d, cap(slice2): %d\n", slice2, len(slice2), cap(slice2))
	fmt.Println()

	slice1 = append(slice1, 500)
	fmt.Println("After append slice1")
	fmt.Printf("slice1: %v, len(slice1): %d, cap(slice1): %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, len(slice2): %d, cap(slice2): %d\n", slice2, len(slice2), cap(slice2))
}
