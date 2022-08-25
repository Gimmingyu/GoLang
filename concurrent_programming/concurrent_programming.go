package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

/*
1000원을 넣고 1000원을 빼는 것을 무한히 반복.
이론상 터지면 안되는데(1000원을 먼저 넣고 나중에 빼니까) 터진다. 왜일까?
이것은 동시성의 문제인데, account라는 인스턴스의 balance에 고루틴들이 동시에 접근한다.
이 때 값을 읽어오는 과정에서 연산을 하고 메모리에 값을 저장하기 전에 잘못된 값을 저장하는 경우가 생길 수 있다.
이러한 문제를 해결하기 위해서는 하나의 자원에 여러 고루틴이 접근할 수 없게 해야한다. ==> Mutex, Semaphore !!
*/
func main() {
	var wg sync.WaitGroup

	account := &Account{10}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}
