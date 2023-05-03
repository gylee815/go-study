package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter...")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Eneter the number")
			stdin.ReadString('\n')
			continue
		} else {
			fmt.Printf("You have enetered %d\n", num)
		}

		if num%2 == 0 {
			break
		}
	}

	fmt.Println("exist for loop")
}
