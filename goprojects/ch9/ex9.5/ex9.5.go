package main

import "fmt"

func hasRichFriend() bool {
	return true
}

func getFriendsCount() int {
	return 3
}

func main() {
	price := 35000

	if price >= 50000 {
		if hasRichFriend() {
			fmt.Println("신발끈 묶는다")
		} else {
			fmt.Println("돈을 나눠 낸다.")
		}
	} else if price >= 30000 {
		if getFriendsCount() > 3 {
			fmt.Println("신발끈 묶는다")
		} else {
			fmt.Println("돈을 나눠 낸다")
		}
	} else {
		fmt.Println("내가 낸다")
	}
}
