package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 6 {
			break
		} else {
			fmt.Printf("6 * %d = %d\n", i, 6*i)
		}
		time.Sleep(time.Second)
	}
}
