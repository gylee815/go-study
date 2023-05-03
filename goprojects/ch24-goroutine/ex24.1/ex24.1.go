package main

import (
	"fmt"
	"time"
)

func PrintAlphabet() {
	for _, v := range "abcdef" {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintAlphabet()
	go PrintNumber()

	time.Sleep(4 * time.Second)
}
