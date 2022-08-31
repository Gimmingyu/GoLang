package main

import (
	"context"
)

type Publisher struct {
	ctx         context.Context
	subscribeCh chan chan<- string
	publishCh   chan string
	subscribers []chan<- string
}

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx: ctx,
		// 채널의 채널... 채널이 들어가는 채널인데, chan<- 는 일방향 채널이다. string 데이터를 넣을 수만 있는 채널이라는 뜻이다
		// <-chan 이라면 readOnly 채널, chan<- 는 writeOnly 채널
		// 아래 예시는 넣기만 하고싶다는 뜻
		subscribeCh: make(chan chan<- string),
		publishCh:   make(chan string),
		subscribers: make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	// 여기서 바로 append 해주지 않는 이유는 slice가 thread safe 하지 않기 때문.
	// 여기서 append 해주려면 lock을 걸어야 한다
	p.subscribeCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscribeCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, subscriber := range p.subscribers {
				subscriber <- msg
			}
		case <-p.ctx.Done():
			wg.Done()
			return
		}
	}
}
