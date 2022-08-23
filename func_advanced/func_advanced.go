package main

import (
	"fmt"
	"os"
)

func sum(numbers ...int) int {
	// ...int == []int
	sum := 0
	fmt.Printf("numbers 타입: %T\n", numbers)
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func every_type(args ...interface{}) {
	for _, arg := range args {
		// ...
		fmt.Println(arg)
	}
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int { return a * b }

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

type OpFn func(int, int) int

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("Value Loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop2")
	for i := 0; i < 3; i++ {
		// Loop마다 새로운 v가 생성되고 할당
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	pl := fmt.Println
	fmt.Printf("type of pl = %T\naddress of pl = %p\n", pl, pl)
	num := sum(1, 2, 3, 4, 5)
	//  == num := sum([]int{1, 2, 3,4, 5})
	pl(num)
	pl(sum(10, 20))

	// defer 사용법
	// defer == 함수 종료전에 실행을 보장
	// 함수가 return 되기 전에 호출된다.
	// 미리 써도, 종료전에 실행된다.
	// 주로 OS 자원을 반납해야 할때 사용
	f, err := os.Create("test.txt")
	if err != nil {
		pl("Failed to create file")
		return
	}

	// Stack의 순서로 호출된다. 밑에서부터 위로
	defer pl("반드시 호출된다.")
	defer f.Close()
	defer pl("파일을 닫았다")

	pl("파일에 Hello world 쓰기")
	fmt.Fprintln(f, "Hello world")

	// 함수 타입 변수
	// 별칭타입 사용 예시
	var operator OpFn
	operator = getOperator("+")

	var result = operator(3, 4)
	pl(result)
	operator = getOperator("*")
	result = operator(3, 4)
	pl(result)

	// 함수 리터럴 (람다)
	// 함수 리터럴은 내부 상태를 가질 수 있다. 매우 유용하게 쓰일 수 있음 !!
	//
	example := func(a, b int) int {
		return a + b
	}
	pl(example(5, 10))

	// 내부상태 예시
	// 캡쳐는 값복사가 아닌 레퍼런스 복사
	// 포인터가 복사된다
	i := 0
	example2 := func() {
		// Parameter가 없는데 변수 사용 가능
		// i의 포인터를 복사됐다
		i += 10
	}
	i++        // i = 1
	example2() // i = 11
	pl(i)

	CaptureLoop()
	CaptureLoop2()

}
