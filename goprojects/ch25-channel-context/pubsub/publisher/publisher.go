package publisher

import (
	"context"
	"sync"
)

type Publisher struct {
	Ctx         context.Context
	SubscribeCh chan chan<- string
	PublishCh   chan string
	Subscribers []chan<- string
	Wg          *sync.WaitGroup
}

func NewPublisher(ctx context.Context, wg *sync.WaitGroup) *Publisher {
	return &Publisher{
		Ctx:         ctx,
		SubscribeCh: make(chan chan<- string),
		PublishCh:   make(chan string),
		Subscribers: make([]chan<- string, 0),
		Wg:          wg,
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.SubscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.PublishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.SubscribeCh:
			p.Subscribers = append(p.Subscribers, sub)
		case msg := <-p.PublishCh:
			for _, subscriber := range p.Subscribers {
				subscriber <- msg
			}
		case <-p.Ctx.Done():
			p.Wg.Done()
			return
		}
	}
}
