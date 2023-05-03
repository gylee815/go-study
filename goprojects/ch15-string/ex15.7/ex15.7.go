package main

import "fmt"

func main() {
	str := "Hello 월드"
	iterate_byte(str)

}

func iterate_len(val string) {
	for i := 0; i < len(val); i++ {
		fmt.Printf("Type: %T, integer value: %d, char: %c\n", val[i], val[i], val[i])
	}

}

func iterate_range(val string) {
	for _, v := range val {
		fmt.Printf("Type: %T, integer value: %d, char: %c\n", v, v, v)
	}
}

func iterate_rune(val string) {
	arr := []rune(val)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type: %T, integer value: %d, char: %c\n", arr[i], arr[i], arr[i])
	}
}

func iterate_byte(val string) {
	arr := []byte(val)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type: %T, integer value: %d, char: %c\n", arr[i], arr[i], arr[i])
	}
}
