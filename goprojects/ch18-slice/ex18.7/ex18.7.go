package main

import (
	"fmt"
)

func main() {
	// slice1, slice2, slice3 모두 같은 array를 가리키고있는 slice
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:5]
	slice3 := append(slice2, 500, 600)

	// slice4는 slice cap을 벋어 났기 때문에 append시 새로운 slice를 반환하므로 slice4는 다른 array를 가리키고 있다
	slice4 := append(slice3, 700)

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))
}
