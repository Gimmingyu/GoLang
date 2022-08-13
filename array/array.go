package main

import (
	"fmt"
	"strconv"
)

func main() {
	var t [5]int32 = [5]int32{1, 2, 3, 4, 5}
	pl := fmt.Println

	for i := 0; i < 5; i++ {
		pl(t[i])
	}

	var nums [5]int

	nums = [5]int{1, 2, 3, 4, 5}
	pl(nums)

	var temps [5]string = [5]string{"1", "2"}
	for i := 0; i < 5; i++ {
		pl(strconv.Atoi(temps[i]))
	} // nil nil invalid invalid invalid
	// Runtime error는 아님

	var s = [5]int{1: 10, 3: 30}
	pl(s)

	x := [...]int{10, 20, 30}
	pl(x)

	y := []int{20, 30}
	pl(y)

	var z []int = []int{40, 30}
	pl(z)
	// 주의할 점
	// [...]int {10, 20, 30} != []int {10, 20, 30}
	// 둘은 완전히 다르다. 전자는 길이가 나중에 정해지는 배열.
	// 후자는 슬라이스다. 슬라이스는 동적배열로 길이가 늘어날수도 줄어들수도 있다.

	z[1] = 300

	pl(z)

	// enumerate ??
	// 단, unused variable이 없도록 해주어야함
	for i, v := range nums {
		pl(i, v)
	}

	// 이거는 괜찮음, _ == 빈칸 지시자
	for _, v := range nums {
		// pl(_, v) 이건 안됨
		pl(v)
	}

	for i := range nums {
		pl(i)
	}

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	// a = b 다른 타입이라서 안됨
	var c [2][5]int = [2][5]int{a, b}
	pl(c)
}
