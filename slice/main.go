package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

// Slice 구조체
//
//	type SliceHeader struct {
//		Data uintptr	// 실제 배열을 가리키는 포인터
//		Len  int		// 요소의 길이
//		Cap  int		// 실제 배열의 길이
//	}
//
// 슬라이스는 배열을 가리키는 포인터 타입이다.
func changeArray(array2 [5]int) {
	// array2 는 call by value가 적용되어서 array가 복사된 객체일 뿐 array가 아님
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	// slice 자체가 배열을 가리키는 포인터라서 변경하면 slice1에도 적용이 된다.
	slice2[2] = 200
}

func main() {
	pl := fmt.Println
	// array := [5]int{1, 2, 3, 4, 5}
	// slice := []int{1, 2, 3, 4, 5}

	// changeArray(array) // 바뀌지않음 ...
	// changeSlice(slice)
	// pl(array)
	// pl(slice)

	// // slice는 go에서 제공하는 동적 배열타입이다.
	// // size는 쓰지 않고 타입만 쓴다.
	// slice1 := []int{1, 2, 3}
	// // var double [][]int

	// slice2 := []int{1, 5: 2, 10: 3}

	// if len(slice1) == 0 {
	// 	pl(slice1)
	// }
	// slice1[1] = 10
	// pl(slice1)
	// fmt.Println(slice2)
	// pl(slice2)

	// /// slice와 map 타입을 만들때 사용 가능한 make
	// var slice3 = make([]int, 3)
	// pl(slice3)
	// slice3[1] = 5
	// pl(slice3)

	// slice4 := make([]int, 10)
	// pl(slice4)

	// // append는 슬라이스에 요소를 추가한 새로운 슬라이스를 반환.
	// // 기존 슬라이스가 바뀔수도 있고 아닐수도 있다.
	// // 빈공간이 충분한가?
	// // yes면 빈공간에 요소추가 (기존 슬라이스)
	// // no면 새로운 배열을 만들고 복사 후 요소추가
	// // 기존 배열을 가리키는 포인터와는 전혀 달라진다.
	// newSlice := []int{1, 2, 3}
	// addSlice := append(newSlice, 4)
	// pl(newSlice, addSlice)

	sc := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pl("sc ", sc)

	// append 이용한 복사(내부적으로 for문 이용해서 도는 것과 같다)
	scTemp := append([]int{}, sc...)
	pl("sc_temp ", scTemp)
	// copy를 이용한 복사
	copied := make([]int, len(sc))
	copy(copied, sc)
	pl("copied ", copied)

	// append를 사용해서 요소 삽입
	idx := 2
	apTest1 := []int{1, 2, 3, 4, 5}
	pl("apTest1 ", apTest1)
	apTest2 := append(apTest1[:idx], apTest1[idx+1:]...)
	pl("apTest2 ", apTest2)

	// copy를 이용해서 요소 삽입
	apTest1 = []int{1, 2, 3, 4, 5}
	pl("apTest1 ", apTest1)
	copy(apTest1[idx:], apTest1[idx+1:])
	pl("apTest2 ", apTest2)

	s := []Student{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}

	sort.Sort(Students(s))
	pl(s)
}

/*
* var array = [...]int{1, 2, 3} // 배열 선언
* var slice = []int{1, 2, 3} // 슬라이스 선언
 */
