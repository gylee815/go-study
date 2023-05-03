package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d th job started\n", j.index)
	time.Sleep(time.Millisecond * 3000)
	fmt.Printf("%d th job finished!! Result: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// job := jobList[i]
		go func(i int) {
			// job.Do()
			jobList[i].Do()
			fmt.Println("wati another 3s...")
			time.Sleep(time.Second * 3)
			wg.Done()
			// fmt.Println("wati another 3s...")
			// time.Sleep(time.Second * 3)
		}(i)
	}
	fmt.Printf("go routine started...\n")
	wg.Wait()
	fmt.Printf("go routine end...\n")
}
