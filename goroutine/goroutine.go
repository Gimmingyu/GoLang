package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		fmt.Printf("%c ", v)
	}
	wg.Done()
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go PrintHangul()
	go PrintNumbers()
	wg.Wait()

}
