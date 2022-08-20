package main

import "fmt"

func main() {
	// slice는 go에서 제공하는 동적 배열타입이다.
	// size는 쓰지 않고 타입만 쓴다.
	slice := []int{1, 2, 3}
	// var double [][]int

	sliec2 := []int{1, 5: 2, 10: 3}

	pl := fmt.Println
	if len(slice) == 0 {
		pl(slice)
	}
	slice[1] = 10
	pl(slice)
	pl(slice2)

}

/*
* var array = [...]int{1, 2, 3} // 배열 선언
* var slice = []int{1, 2, 3} // 슬라이스 선언
 */
