package main

import (
	"fmt"
)

func main() {
	printThreeDimensionalArray()
}

func printTwoDimensionalArray() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a {
		fmt.Printf("{ ")
		for _, v := range arr {
			fmt.Printf("%d, ", v)
		}
		fmt.Printf("}\n")
	}
}

func printThreeDimensionalArray() {
	a := [2][2][5]int{
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
		},
		{
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
		},
	}

	for _, second_arr := range a {
		fmt.Printf("{\n")
		for _, arr := range second_arr {
			fmt.Printf("{ ")
			for _, v := range arr {
				fmt.Printf("%d, ", v)
			}
			fmt.Printf("}\n")
		}
		fmt.Printf("},\n")
	}
}
