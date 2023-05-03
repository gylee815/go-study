package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		bufferdData, _ := stdin.ReadString('\n')
		fmt.Println(bufferdData)
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		bufferdData, _ := stdin.ReadString('\n')
		fmt.Println(bufferdData)
	} else {
		fmt.Println(n, a, b)
	}
}
