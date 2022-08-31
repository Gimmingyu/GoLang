package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover :", r)
		}
	}()
	// 2칸짜리 자리를 만들어주게되면, 자리를 설정하지 않았을 때와 다르게 Never Print가 실행된다.
	// 자리를 만들어주게 되면 채널에 9라는 값을 놓고 다음으로 넘어간다.
	// 자리를 미리 설정해두지 않으면 9라는 값을 누군가 가져갈 때까지 기다린다.
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
	/*
		채널에서 무한히 뽑아내려고 한다. 결국 고루틴이 종료되지 않으면서 deadlock 발생
		채널을 닫아주지 않아서 무한 대기를 하는 고루틴을 좀비 고루틴 또는 고루틴 릭이라고 한다.
	*/
	for n := range ch {
		fmt.Println("Square:", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
