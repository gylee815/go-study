package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)

	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square %d = %d\n", n, n*n)
		time.Sleep(time.Second * 2)
	}
	wg.Done()
}
