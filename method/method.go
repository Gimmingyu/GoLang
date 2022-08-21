package main

import "fmt"

/**
메서드는 타입에 속한 함수이다.

함수는 독립적이나, 메서드는 타입에 종속된다.


func (r Rabbit) info() int {
	return r.width * r.height
}

r Rabbit	= 리시버
info()		= 메서드명

다른 언어에서는 클래스 내에서 메서드를 정의...
고에서는 타입 정의 따로, 메서드 정의 따로, 대신 리시버 이름을 적어서 어떤 타입에 속하는지 정의
*/

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 그냥 함수
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	pl := fmt.Println

	pl("Start")

	// a := &account{100}
	// pl(a)
	// withdrawFunc(a, 30)
	// pl(a)
	// a.withdrawMethod(30)
	// pl(a)

	// var a myInt
	// a = a.Add(10)
	// pl(a)

	var mainA *account = &account{100, "joe", "Park"}
	mainA.withdrawPointer(30)
	pl(*mainA)

	mainA.withdrawValue(20)
	pl(*mainA)
}
