package main

import "fmt"

type Data struct {
	Number int
}

func NewData(num int) *Data {
	var newData = Data{num}
	return &newData
}

func main() {
	pl := fmt.Println
	var a int
	var p *int

	p = &a
	pl(a, p)
	a = 10
	pl(a, p)

	pl(*p)

	var str string = "hello world"
	var strPtr *string = &str

	pl(*strPtr)
	pl(str)
	for i, v := range *strPtr {
		pl(i, string(v))
	}

	var structure *Data = &Data{30}
	pl(*structure)

	p1 := &Data{}
	var p2 = new(Data)

	pl(p1, p2)

	// Escape Analazing ... 탈출하는 인스턴스는 힙에 생성시킨다.
	p3 := NewData(10)
	pl(p3)
}
