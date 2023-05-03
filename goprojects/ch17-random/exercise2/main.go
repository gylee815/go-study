package main

import (
	"bufio"
	"fmt"
	slot "goprojects/ch17-exercise2/slot"
	"os"
)

func main() {
	var m [3][10]int
	slot.InitSlotMachine(&m)
	fmt.Println(m)

	// rst := SlotResult(&slot)
	// fmt.Println(*rst)

	var stdin = bufio.NewReader(os.Stdin)
	var isContinue string
	myMoney := 10000
	trial := 1

	for myMoney > 1000 {
		fmt.Printf("trial: %d\n", trial)

		if trial > 1 {
			for {
				fmt.Printf("Continue? (y/n): ")
				val, err := InputIntValue(stdin)
				if err != nil {
					fmt.Println(err)
					continue
				}

				if val == "y" || val == "n" {
					isContinue = val
					break
				} else {
					fmt.Println(val)
				}
			}
		}

		if isContinue == "n" {
			break
		}

		myMoney -= 2000

		mySlot := slot.RunSlotMachine(&m)
		fmt.Println(mySlot)

		switch slot.SlotSort(&mySlot) {
		case slot.Jackpot:
			fmt.Println("Jackpot!!!")
			myMoney += 500000
		case slot.Equalone:
			fmt.Println("Equal two values")
			myMoney += 4000
		case slot.Equalzero:
			fmt.Println("Oops... sorry. Try again")
		}
		trial += 1

		fmt.Printf("remain money: %d\n\n", myMoney)
	}

	if isContinue != "n" {
		fmt.Println("You lose all your money...")
	}
	fmt.Println("See you next time Bye :)")
}

func InputIntValue(stdin *bufio.Reader) (string, error) {
	var val string
	_, err := fmt.Scanln(&val)
	if err != nil {
		stdin.ReadString('\n')
	}

	if val != "y" && val != "n" && val != "Y" && val != "N" {
		return "Wrong value inputed", err
	}

	return val, err
}

// func InputIntValueByte(stdin *bufio.Reader) (*[]byte, error) {
// 	var val string
// 	var rst []byte
// 	_, err := fmt.Scanln(&val)
// 	if err != nil {
// 		stdin.ReadString('\n')
// 	}

// 	if val != "y" && val != "n" && val != "Y" && val != "N" {
// 		rst = []byte("Wrong value inputed")
// 		return &rst, err
// 	}

// 	rst = []byte(val)
// 	return &rst, err
// }

// func CompareSlot(slot *[3]int) {
// 	for i, v := range slot {

// 	}
// }
