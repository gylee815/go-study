package ex1

import "fmt"

type Product struct {
	Name         string
	Price        int
	ReviewSource float64
}

func PrintStruct(p Product) {
	fmt.Println(p)
}
