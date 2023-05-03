package main

import "fmt"

func f() {
	fmt.Println("f() function started")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %s\n", r)
		}
	}()

	g()
	fmt.Println("f() function finished")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
}

func h(a, b int) int {
	if b == 0 {
		panic("cannot devide with 0")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("program continue")
}
