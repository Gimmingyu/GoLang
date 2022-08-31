package main

import (
	"fmt"
	"sync"
	"time"
)

/*
select는 여러 채널에서 동시에 데이터를 기다릴 때 사용한다
select는 for문 예시와 다르게 아래 케이스 중에서 하나가 실행되면 빠져나오게 된다.
select {
case n := <- ch1:

	...

case n2 := <- ch2:

	...

case ...
}

for문은 하나의 채널만 기다리고 계속 반복
for n := range ch {

}
*/
var wg sync.WaitGroup

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover :", r)
		}
	}()
	ch := make(chan int)
	wg.Add(1)
	go square(ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
	fmt.Println("Never print")
}

func square(ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("Square:", n*n)
			time.Sleep(time.Second)
		}
	}
}
