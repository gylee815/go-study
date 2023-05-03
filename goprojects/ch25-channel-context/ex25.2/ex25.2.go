package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int, 2)
	go square(ch, &wg)

	go func(wg *sync.WaitGroup, ch chan int) {
		for {
			var val int
			fmt.Scanln(&val)

			if val == 0 {
				close(ch)
				wg.Done()
				return
			}

			ch <- val
			// time.Sleep(time.Second * 2)
		}
	}(&wg, ch)

	// ch <- val
	// time.Sleep(time.Second * 2)
	// ch <- 8
	// time.Sleep(time.Second * 2)
	// ch <- 7
	// time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("Never print")
}

func square(ch chan int, wg *sync.WaitGroup) {
	tmp := 0
	for n := range ch {
		// n := <-ch
		fmt.Println(tmp + n)
		tmp = n
		// fmt.Println("channel got input value \n", n)
		// time.Sleep(time.Second)
		// fmt.Printf("%d * %d = %d\n", n, n, n*n)
		// wg.Done()
	}
	wg.Done()
}
