package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)
	go square(&wg, ch)
	time.Sleep(time.Second * 2)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch
	fmt.Println("channel got input value")
	time.Sleep(time.Second)
	fmt.Printf("%d * %d = %d", n, n, n*n)
	wg.Done()
}
