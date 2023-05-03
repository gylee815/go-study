package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "data.txt"

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprint(file, line)
	return err
}

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Printf("Failed to write file!! Error: %w\n", err)
			return
		} else {
			line, err = ReadFile(filename)
			if err != nil {
				fmt.Printf("Failed to read file!! Error: %w\n", err)
				return
			}
		}
	}
	fmt.Printf("File: %s\n", line)
}
