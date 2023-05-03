package main

import (
	"fmt"
	"strconv"
)

func printNo(n int) string {
	if n == 0 {
		return "returned" + strconv.Itoa(n)
	}

	fmt.Println(n)
	fmt.Println(printNo(n - 1))
	fmt.Println("After", n)

	return "returned" + strconv.Itoa(n)
}

func main() {
	fmt.Println(printNo(3))
}
