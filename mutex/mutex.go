package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	account := &Account{10}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
				fmt.Println(account.Balance)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000

}

/*
** 뮤텍스의 문제점 **
동시성 프로그래밍으로 인한 성능 향상을 얻을 수  없다.
심지어 과도한 Locking으로 성능이 하락되기도 한다.
Lock을 획득하는 과정, Lock을 푸는 과정에서 이미 자원을 많이 쓴다.
또한 고루틴을 완전히 멈추게 만드는 데드락 문제도 발생할 수 있다.
*/
