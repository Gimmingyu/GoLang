package main

import "fmt"

/*
	Go는 SEH(Structured Error Handling)을 지원하지 않는다. (try-catch, try-except)
	SEH를 사용하면 성능적인 문제 + 에러를 먹어버리는 문제 (에러 처리를 등한시하게 된다)
	보통 에러 처리를 귀찮아 하고 코드를 지저분하게 만든다고 생각하지만 에러 처리는 매우 중요하다
	에러를 반환하는 함수에서 반환되는 에러를 제대로 처리해야한다 (_ 빈칸 지시자로 무시하는 것은 매우 안좋은 습관)
*/

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
