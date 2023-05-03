package subscriber

import (
	"context"
	"fmt"
	publisher "goprojects/pubsub/publisher"
	"sync"
)

type Subscriber struct {
	Ctx   context.Context
	Name  string
	MsgCh chan string
	Wg    *sync.WaitGroup
}

func NewSubscriber(name string, ctx context.Context, wg *sync.WaitGroup) *Subscriber {
	return &Subscriber{
		Ctx:   ctx,
		Name:  name,
		MsgCh: make(chan string),
		Wg:    wg,
	}
}

func (s *Subscriber) Subscribe(pub *publisher.Publisher) {
	pub.Subscribe(s.MsgCh)
}

func (s *Subscriber) Update() {
	for {
		select {
		case msg := <-s.MsgCh:
			fmt.Printf("%s got Message: %s\n", s.Name, msg)
		case <-s.Ctx.Done():
			s.Wg.Done()
			return
		}
	}
}
