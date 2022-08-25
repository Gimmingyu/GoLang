package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 다음 단어를 읽어서 숫자로 변환하여 반환합니다
// 변환된 숫자, 읽은 글자수, 에러를 반환합니다
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to integer, word: %s, err: %w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
		}
	}
}

func MultipleFromString(str string) (int, error) {
	// 어떤 데이터를 어떤 규칙으로 가져올 때 편하다.
	// 인자로 io reader를 받는다.
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d err: %w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d err: %w", pos, err)
	}
	return a * b, nil
}
func main() {
	readEq("123 3")
	readEq("123 abc")
}
