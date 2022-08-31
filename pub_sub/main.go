package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
발행(Publisher)/구독(Subscriber) 패턴 구현하기
옵저버 패턴과 거의 유사(같다고 봐도 무방)

옵저버 패턴이 기본적으로 동기적으로 작동해서 나온 패턴이라고 한다.
*/

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(4)
	publisher := NewPublisher(ctx)
	subscriber1 := NewSubscriber("AAA", ctx)
	subscriber2 := NewSubscriber("BBB", ctx)

	go publisher.Update()

	subscriber1.Subscribe(publisher)
	subscriber2.Subscribe(publisher)

	go subscriber1.Update()
	go subscriber2.Update()

	go func() {
		tick := time.Tick(time.Second / 10)
		for {
			select {
			case <-tick:
				publisher.Publish("Hello world")
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
