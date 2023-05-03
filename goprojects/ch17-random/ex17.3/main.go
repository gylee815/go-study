package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var stdin = bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())
	ranNum := rand.Intn(100)
	trial := 1

	for {
		fmt.Print("Enter number: ")
		inNum, err := InputIntValue(stdin)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if GuessNum(inNum, ranNum, &trial) {
			break
		} else {
			trial++
			fmt.Println("You have entered wrong number!")
			continue
		}
	}

}

func InputIntValue(stdin *bufio.Reader) (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func GuessNum(inNum int, ranNum int, trial *int) bool {
	if ranNum > inNum {
		fmt.Println("Up")
	} else if ranNum < inNum {
		fmt.Println("Down")
	} else {
		fmt.Printf("You got answer! ranNum was %d, You have tried: %d\n", ranNum, *trial)
		return true
	}
	return false
}
