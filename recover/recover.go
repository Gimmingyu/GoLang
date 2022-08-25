package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		// 리커버를 시도, 결과가 있다면 복구
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g()

	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("%d / %d = %d\n", 9, 3, h(9, 3))
	fmt.Printf("%d / %d = %d\n", 9, 0, h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}
