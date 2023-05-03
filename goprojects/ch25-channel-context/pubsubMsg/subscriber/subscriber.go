package subscriber

import (
	"context"
	"fmt"
	broker "goprojects/pubsubMsg/broker"
	"sync"
)

type Subscriber struct {
	Ctx   context.Context
	Name  string
	MsgCh chan string
	Wg    *sync.WaitGroup
}

func NewSubscriber(ctx context.Context, wg *sync.WaitGroup, name string) *Subscriber {
	return &Subscriber{
		Ctx:   ctx,
		Name:  name,
		MsgCh: make(chan string),
		Wg:    wg,
	}
}

func (s *Subscriber) Subscribe(b *broker.Broker) {
	b.SubscribeCh <- s.MsgCh
	// p.Subscribe(s.MsgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		case msg := <-s.MsgCh:
			fmt.Printf("%s got Message: %s !!\n", s.Name, msg)
		case <-s.Ctx.Done():
			s.Wg.Done()
			return
		}
	}
}
