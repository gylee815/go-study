package main

import "fmt"

func devide(a, b int) (int, bool) {
	if a%b == 0 {
		return a / b, true
	} else {
		return 0, false
	}
}

func main() {
	if num, success := devide(3, 2); success {
		fmt.Printf("Devide %v, returned value %d\n", success, num)
	} else {
		fmt.Printf("Devide %v, returned value %d\n", success, num)
	}
}
