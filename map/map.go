package main

import "fmt"

/*
맵은 키와 값 형태로 데이터를 저장하는 자료구조이다.
언어에 따라 딕셔너리, 해쉬테이블, 해쉬맵 등으로 불린다.
map[key]value  == key에는 key 타입이 들어가고, value에는 값타입이 들어간다.
map[string]int == "aaa" => 3 과 같은 방식 ...
*/

// MakeMap ...
func MakeMap() map[string]string {
	var newMap = make(map[string]string)
	return newMap
}

func main() {
	m := MakeMap()

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "부산시 사하구"
	m["최번개"] = "전주시 덕진구"

	for k, v := range m {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	println("=================================================================")
	m["최번개"] = "청주시 상당구"
	for k, v := range m {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	// 요소 추가
	m["김민규"] = "성남시 수정구"
	// 요소 삭제
	delete(m, "김민규")
	// 존재 여부 확인
	if v, ok := m["김민규"]; ok {
		fmt.Println(v)
	}
}
