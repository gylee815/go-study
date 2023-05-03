package main

import (
	"context"
	"fmt"
	pub "goprojects/pubsub/publisher"
	sub "goprojects/pubsub/subscriber"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := pub.NewPublisher(ctx, &wg)
	subscriber1 := sub.NewSubscriber("AAA", ctx, &wg)
	subscriber2 := sub.NewSubscriber("BBB", ctx, &wg)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello Message!!")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()

	fmt.Scanln()
	cancel()

	wg.Wait()
}
