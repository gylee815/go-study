package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("Failt to open or create file, reason: %v", err)
		return
	}

	// defer는 윗줄부터 stack으로 쌓인다. 그러므로 출력 결과는 아래서부터 출력된다.
	defer fmt.Println("file is closed sucessfully!") //3
	defer f.Close()                                  //2
	defer fmt.Println("file close!")                 //1

	fmt.Fprintf(f, "Type of f: %T, Size of f: %v\n", f, unsafe.Sizeof(f))
}
