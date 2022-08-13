package main

import "fmt"

func main() {
	var typeofVar int32 = 12

	fmt.Printf("Type = %T\n", typeofVar)

	var a = 123

	fmt.Printf("%v\n", a)
	fmt.Printf("%d\n", a)

	var str string

	fmt.Scan(&str)
	println(str)

	fmt.Scanln(&str)
	println(str)
}
