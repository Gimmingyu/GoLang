package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello, Variable!"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var u8 uint8 = 255

	fmt.Print(u8)

	var u16 uint16 = 65535

	fmt.Println(u16)

	var u32 uint32 = 4294967295

	fmt.Println(u32)

	var i32 int32 = 2147483647

	fmt.Println(i32)

	// and etc ..

}
