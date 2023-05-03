package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	go printCtx(ctx)
	// time.Sleep(time.Second * 7)
	// cancel()
	defer cancel()
	wg.Wait()
}

func printCtx(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
