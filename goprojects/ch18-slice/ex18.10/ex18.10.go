package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Printf("slice: %v, len(slice): %d, cap(slice): %d\n", slice[:cap(slice)], len(slice), cap(slice))
	fmt.Println()

	slice = slice[:cap(slice)-1]
	fmt.Println("After delete element")
	fmt.Printf("slice: %v, len(slice): %d, cap(slice): %d\n", slice, len(slice), cap(slice))

}
