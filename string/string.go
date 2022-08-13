package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func Upper(str string) string {
	var ret string
	for _, c := range str {
		if c >= 'a'**c <= 'z' {
			ret += string('A' - (c - 'a'))
		} else {
			ret += string(c)
		}
	}
	return ret
}

func ToUpper(str string) string {

	// 내부에서 슬라이스를 가지고 있다.
	// 글자마다 reallocate가 아니라, 미리 공간을 좀 할당해놓고 넣어준다.
	// 복사비용도 아끼고 메모리 공간도 아낄 수 있다
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	pl := fmt.Println

	pl("Start")

	start := "Start"

	pl(len(start))
	for i, v := range start {
		pl(i, v)
	}
	for i, v := range start {
		pl(i, string(v))
	}

	for i := 0; i < len(start); i++ {
		pl(start[i])
	}

	arr := []rune(start)
	// 슬라이스는 동적 배열
	// rune ?? int32 별칭타입
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Type: %T, Value: %d, String Value: %c\n", arr[i], arr[i], arr[i])
	}

	str1 := "Hello, world!"
	str2 := str1

	// &str1 을 Pointer 타입으로 변환 ... == string Pointer => unsafe.Pointer로 변환
	// StringHeader는 ?? Data + Len 으로 이루어진 구조체
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	// 똑같은 메모리 공간을 가리키고 있고, 똑같은 길이를 가지고 있다.
	pl(stringHeader1)
	pl(stringHeader2)

	var str string = "Hello world"
	var slice []byte = []byte(str)
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	// 주소를 가진다, string과 slice가 각각 다른 위치에 있다.. 타입 변환을 하면서 문자열이 복제 된다.
	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)

	var newStr = ToUpper(str)
	pl(newStr)
}
