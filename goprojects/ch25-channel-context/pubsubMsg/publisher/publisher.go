package publisher

import (
	"context"
	"fmt"
	broker "goprojects/pubsubMsg/broker"
	"sync"
)

type Publisher struct {
	Ctx    context.Context
	Name   string
	PubMsg chan string
	// Msg    string
	Cancel context.CancelFunc
	Wg     *sync.WaitGroup
}

func NewPublisher(ctx context.Context, wg *sync.WaitGroup, name string, cancel context.CancelFunc) *Publisher {
	return &Publisher{
		Ctx:    ctx,
		Wg:     wg,
		Name:   name,
		PubMsg: make(chan string),
		// Msg:    "",
		Cancel: cancel,
	}
}

func (p *Publisher) Publish() {
	// if len(p.PubMsg) == 0 {
	// 	var msg string
	// 	fmt.Scanln(&msg)
	// 	p.PubMsg <- msg
	// }

	for {
		var msg string
		fmt.Scanln(&msg)
		p.PubMsg <- msg

		if msg == "x" {
			p.Wg.Done()
			return
		}
	}
}

func (p *Publisher) Update(b *broker.Broker) {
	for {
		select {
		case msg := <-p.PubMsg:
			if msg == "x" {
				close(p.PubMsg)
				// return
				p.Wg.Done()
				p.Cancel()
				fmt.Println("finished")
				return
			} else {
				b.PublishCh <- msg
			}
		case <-p.Ctx.Done():
			// p.Wg.Done()
			return
		}
	}
}
