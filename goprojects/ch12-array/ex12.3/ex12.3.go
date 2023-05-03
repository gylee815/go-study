package main

import (
	"fmt"
)

func main() {
	search_2()
}

func search_1() [5]int {
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Printf("num: %d\n", nums[i])
	}

	return nums
}

func search_2() {
	// nums := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	str := "giyeanlee"
	for i, v := range str {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}
}
