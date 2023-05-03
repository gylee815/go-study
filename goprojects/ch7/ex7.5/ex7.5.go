package main

import "fmt"

func Devide(a, b int) (rst int, suc bool) {
	if b == 0 {
		rst = 0
		suc = false
		return
	} else {
		rst = a / b
		suc = true
		return
	}
}

func main() {
	x, success := Devide(9, 3)
	fmt.Printf("result : %d, %v\n", x, success)

	y, success := Devide(9, 0)
	fmt.Printf("result : %d, %v\n", y, success)
}
