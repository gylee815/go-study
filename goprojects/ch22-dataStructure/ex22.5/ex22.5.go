package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["lee"] = "Seoul Gangnam"
	m["kim"] = "Seoul Gangseo"
	m["park"] = "Daejeon Seogu"
	m["kang"] = "Daejeon Donggu"

	m["kang"] = "Incheon Seogu"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
