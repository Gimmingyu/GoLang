package main

import (
	"fmt"
	"strconv"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	a := 10
	b := 20
	c := Add(a, Add(b, a))
	fmt.Println(c)

	str := "10000"
	d, _ := strconv.Atoi(str)
	fmt.Println(d)

	fibo := fibonazzi(10)
	println(fibo)
}

func fibonazzi(a int) int {
	if a == 1 {
		return 1
	} else if a == 2 {
		return 1
	} else if a == 0 {
		return 0
	}
	return fibonazzi(a-1) + fibonazzi(a-2)
}
