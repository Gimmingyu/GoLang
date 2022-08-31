package main

import (
	"fmt"
	"sync"
)

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
	/*
		넣을거 다 넣고 닫아버리면 좀비 고루틴 및 데드락 방지 가능
		좀비 고루틴이 생기면 시스템 부하가 커지기 때문에 반드시 채널을 받아주는 습관을 들이자
	*/
	close(ch)
	wg.Wait()
	fmt.Println("Never print")
}

func square(ch chan int) {
	for n := range ch {
		fmt.Println("Square:", n*n)
	}
	wg.Done()
}
