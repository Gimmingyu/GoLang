package main

import (
	"fmt"
)

func main() {
	pl := fmt.Println
	for i := 0; i < 10; i++ {
		pl("i =", i)
	}

	i := 0

	for i < 10 {
		pl(i)
		i++
	}
	for i < 100 {
		pl(i)
		i += 10
	}

	// for true {
	// 	pl("무한루프 ON")
	// }

	// for {
	// 	pl("무한루프 ON")
	// }

	// for {
	// 	time.Sleep(time.Second)
	// 	pl(i)
	// 	i++
	// }

	a := 1
	b := 1
OuterFor: // Label 추천하지는 않는다. 강력하지만 위험하다 ( 스택 변수 스코프가 꼬일 수 있음 )
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}

	pl(a, b)
}
