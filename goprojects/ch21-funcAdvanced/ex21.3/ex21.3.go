package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type OpFunc func(int, int) int

func getOperator(op string) OpFunc {
	switch op {
	case "+":
		return add
	case "*":
		return mul
	default:
		return nil
	}
}

func getOperator2(op string) OpFunc {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

func main() {
	op := getOperator("+")
	fmt.Printf("Type of op: %T\n", op)

	rst := op(3, 4)
	fmt.Println(rst)

	op2 := getOperator2("*")
	fmt.Printf("Type of op: %T\n", op2)

	rst2 := op2(3, 4)
	fmt.Println(rst2)
}
