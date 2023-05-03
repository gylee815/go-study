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
	curTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	ctx = context.WithValue(ctx, "current", curTime)
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
			if v := ctx.Value("current"); v != nil {
				cur := v.(time.Time)
				fmt.Printf("duration: %.2f passed \n", time.Since(cur).Seconds())
			}
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}
