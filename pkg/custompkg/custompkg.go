package custompkg

import "fmt"

// 소문자는 외부로 공개되지 않는다.
type Student struct {
	Name string
	Age  int
	// 얘만 못씀
	score int
}

type name struct {
	i int
}

func PrintCustom() {
	pl := fmt.Println
	pl("This is a custom package")
}

func printcustom2() {
	pl := fmt.Println
	pl("This is a custom package")
}
