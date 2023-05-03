package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b, num int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("goroutine num[%d]: sum of %d to %d = %d\n", num, a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go SumAtoB(1, 1000000, i)
	}

	wg.Wait()
}
