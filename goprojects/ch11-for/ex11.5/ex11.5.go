package main

import (
	"fmt"
)

func main() {
	starPyrmid()
}

func starIncrease() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starDecrease() {
	end := 5
	for i := 0; i < end; i++ {
		for j := end; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func starPyrmid() {
	end := 9
	for a, i := 0, 0; i < end; i++ {
		if (i+1)%2 == 0 {
			continue
		}
		for j := 0; j < end; j++ {
			if j >= end/2-a && j <= end/2+a {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		a++
		fmt.Println()
	}
}
