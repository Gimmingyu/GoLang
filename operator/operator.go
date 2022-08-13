package main

import "fmt"

func main() {
	var a = 10
	var b = 34

	var c = a & b
	var d = a | b
	fmt.Printf("a == %d\n", a)
	fmt.Printf("b == %d\n", b)

	a, b = b, a

	fmt.Printf("a == %d\n", a)
	fmt.Printf("b == %d\n", b)
	if a == 10 && b == 34 {
		fmt.Printf("c == %d\n", c)
	} else {
		fmt.Printf("d == %d\n", d)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	var i = 0
	for {
		if i == 10 {
			println("Break !!")
			break
		}
		fmt.Printf("i = %d\n", i)
		i++
	}
}
