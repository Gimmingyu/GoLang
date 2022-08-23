package main

import (
	"fmt"
	"math"
)

// fmt.Errorf(formatter string, ...interface{}) error

// errors.New(text string) error

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("sqrt(x) must be positive value. f: %g", f)
	}
	return math.Sqrt(f), nil
}

func main() {

	// sqrt, err := sqrt(-2)
	// if err != nil {
	// 	fmt.Printf("Error : %v\n", err)
	// }
	// fmt.Printf("Sqrt(-2) = %v\n", sqrt)

	// 가장 중요한 구조..
	// 원하는 에러의 인터페이스를 만들어놓고, 인스턴스를 형변환 시켜서 에러를 핸들링한다.
	err := RegisterAccount("MyId", "MyPw")
	if err != nil {
		// 멋지게 에러 핸들링 하는 법
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len : %d RequireLen: %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입 완료")
	}
}

/*
type error interface {
	Error() string
}
*/
