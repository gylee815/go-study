package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Cannot find file!! ", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
