package main

import "fmt"

func main() {
	const (
		BitFlag1 uint = 1 << iota
		BitFlag2
		BitFlag3
		BitFlag4
		BitFlag5
		BitFlag6
		BitFlag7
		BitFlag8
	)

	fmt.Printf("%08b\n", BitFlag1)
	fmt.Printf("%08b\n", BitFlag2)
	fmt.Printf("%08b\n", BitFlag3)
	fmt.Printf("%08b\n", BitFlag4)
	fmt.Printf("%08b\n", BitFlag5)
	fmt.Printf("%08b\n", BitFlag6)
	fmt.Printf("%08b\n", BitFlag7)
	fmt.Printf("%08b\n", BitFlag8)
}
