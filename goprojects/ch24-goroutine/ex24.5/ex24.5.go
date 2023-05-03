package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go DiningProblem("A", fork, spoon, "fork", "spoon")
	go DiningProblem("B", spoon, fork, "spoon", "fork")
	wg.Wait()
}

func DiningProblem(name string, first, second *sync.Mutex, firstGet, secondGet string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s is ready to eat\n", name)
		first.Lock()
		fmt.Printf("%s get %s\n", name, firstGet)
		second.Lock()
		fmt.Printf("%s get %s\n", name, secondGet)
		fmt.Printf("%s start to eat\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		defer first.Unlock()
		defer second.Unlock()
	}
}
