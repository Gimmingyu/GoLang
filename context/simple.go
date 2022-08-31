package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
컨텍스트 : 작업을 지시할 때 작업 가능시간, 작업 취소 등의 조건을 지시할 수 있는 작업 명세서 역할
*/

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// 컨텍스트의 Background는 기본 컨텍스트를 하나 반환해준다
	// 컨텍스트를 만들 때는 항상 기본 컨텍스트에 덮어 쓰게 되어있다.
	// 기본으로 만들어진 컨텍스트가 ctx가 되고, 취소할 수 있는 WithCancel 함수가 cancel에 담긴다.
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		// 작업이 끝날 때 시그널이 온다. --> cancel에서 받은 시그널이 오는 것
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
