package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[0] = Product{"Nike", 50000}
	m[1] = Product{"Addidas", 60000}
	m[100] = Product{"Reebok", 45000}
	m[200] = Product{"FLA", 35000}
	m[340] = Product{"Puma", 55000}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
