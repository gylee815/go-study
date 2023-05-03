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
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminated := time.After(time.Second * 10)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case t := <-terminated:
			fmt.Printf("Terminated after %d seconds\n", t)
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square %d = %d\n", n, n*n)
			time.Sleep(time.Second)
		}
	}
}
