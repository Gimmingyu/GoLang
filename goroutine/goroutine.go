package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		fmt.Printf("%c \n", v)
	}
	wg.Done()
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d 부터 %d까지 합계는 %d 입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(2)
	go PrintHangul()
	go PrintNumbers()
	wg.Wait()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go sumAtoB(1, 1000000000)
	}

	wg.Wait()
}
