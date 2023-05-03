package broker

import (
	"context"
	"sync"
)

type Broker struct {
	Ctx         context.Context
	SubscribeCh chan chan<- string
	PublishCh   chan string
	Subscribers []chan<- string
	// Publishers  []chan<- string
	Wg *sync.WaitGroup
}

func NewBroker(ctx context.Context, wg *sync.WaitGroup) *Broker {
	return &Broker{
		Ctx:         ctx,
		SubscribeCh: make(chan chan<- string),
		PublishCh:   make(chan string),
		Subscribers: make([]chan<- string, 0),
		// Publishers: make([]chan<- string, 0),
		Wg: wg,
	}
}

// func (p *Publisher) Subscribe(sub chan<- string) {
// 	p.SubscribeCh <- sub
// }

func (b *Broker) Publish(msg string) {
	b.PublishCh <- msg
}

func (b *Broker) Update() {
	for {
		select {
		case sub := <-b.SubscribeCh:
			b.Subscribers = append(b.Subscribers, sub)
		case msg := <-b.PublishCh:
			for _, subscriber := range b.Subscribers {
				subscriber <- msg
			}
		case <-b.Ctx.Done():
			b.Wg.Done()
			return
		}
	}
}
