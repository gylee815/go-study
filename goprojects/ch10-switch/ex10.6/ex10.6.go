package main

import (
	"fmt"
)

func getMyAge() int {
	return 44
}

func main() {
	switch age := getMyAge(); true {
	case age >= 10 && age < 20:
		fmt.Println("Teenage")
	case age >= 20 && age < 30:
		fmt.Println("Twenties")
	case age >= 30 && age < 40:
		fmt.Println("Thirties")
	default:
		fmt.Printf("My age is %d\n", age)
	}

	// fmt.Printf("Age is %d", age)
}
