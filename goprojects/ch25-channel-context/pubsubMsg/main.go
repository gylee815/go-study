package main

import (
	"context"
	broker "goprojects/pubsubMsg/broker"
	pub "goprojects/pubsubMsg/publisher"
	sub "goprojects/pubsubMsg/subscriber"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// inputCh := make(chan string)

	wg.Add(5)

	broker := broker.NewBroker(ctx, &wg)
	publisher := pub.NewPublisher(ctx, &wg, "publisher1", cancel)
	subscriber1 := sub.NewSubscriber(ctx, &wg, "gylee")
	subscriber2 := sub.NewSubscriber(ctx, &wg, "gblee")

	go broker.Update()

	// publisher.Publish(broker)
	subscriber1.Subscribe(broker)
	subscriber2.Subscribe(broker)

	go publisher.Publish()
	go publisher.Update(broker)
	go subscriber1.Update()
	go subscriber2.Update()

	// go inputMsg(&wg, inputCh, cancel)

	// go func(inputCh chan string) {
	// 	// tick := time.Tick(time.Second * 2)
	// 	for {
	// 		select {
	// 		case msg := <-inputCh:
	// 			// fmt.Println(msg)
	// 			if msg != "" {
	// 				// fmt.Println(msg)
	// 				broker.Publish(msg)
	// 			}
	// 		case <-ctx.Done():
	// 			wg.Done()
	// 			return
	// 		}
	// 	}
	// }(inputCh)

	// fmt.Scanln()
	// cancel()

	wg.Wait()
}

// func inputMsg(wg *sync.WaitGroup, inputCh chan string, cancel context.CancelFunc) {
// 	for {
// 		var message string
// 		if len(inputCh) == 0 {
// 			// fmt.Printf("input messge: ")
// 			fmt.Scanln(&message)
// 		}

// 		// fmt.Println()
// 		if message == "x" {
// 			close(inputCh)
// 			cancel()
// 			wg.Done()
// 			return
// 		}
// 		inputCh <- message
// 	}
// }
