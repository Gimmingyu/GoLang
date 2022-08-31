package main

import (
	"log"
	"sync"
	"time"
)

/*
채널, 컨텍스트, 고루틴은 GoLang의 3신기라고 불린다.
동시성 프로그래밍의 3신기로서 핵심이라고 볼 수 있다.

채널은 고루틴끼리 메세지를 전달할 수 있는 메세지 큐이다.
FIFO 형태로 진행되는 큐인데, Thread-safe한 큐이다.

멀티스레드 환경에서 Lock 없이 쓸 수 있다. (미친거 같다)

*/

// MakeChannel : 채널 생성하는 방법. Make()를 이용해서 만들 수 있다.
func MakeChannel() chan string {
	// chan string == 채널이고, 큐에 들어가는 요소의 타입은 string이다
	var messages chan string = make(chan string)
	return messages
}

// PopDataFromChannel : 채널에서 데이터 빼기
func PopDataFromChannel(channel chan string, wg *sync.WaitGroup) {
	var msg string = <-channel
	log.Println(msg)
	time.Sleep(time.Second)
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover Success")
		}
	}()
	c := MakeChannel()
	wg.Add(1)
	// go PushDataToChannel(&wg, c, "This is a test message")
	go PopDataFromChannel(c, &wg)
	c <- "This is a test message2"
	wg.Wait()

}

// PushDataToChannel : 채널에 데이터 넣기
func PushDataToChannel(wg *sync.WaitGroup, ch chan string, data string) {
	ch <- data
	wg.Done()
}
