package main

import (
	"bufio"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	// 어떤 데이터를 어떤 규칙으로 가져올 때 편하다.
	// 인자로 io reader를 받는다.
	scanner := bufio.NewScanner(strings.NewReader(str))
}
func main() {

}
