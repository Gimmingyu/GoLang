package main

import (
	"fmt"
	"os"
)

type Writer func(string)
type WriterInterface interface {
	Write(string)
}

func WriteHello(writer Writer) {
	// 의존성 주입
	writer("Hello world")
}
func WriteHello2(writer WriterInterface) {
	// 의존성 주입
	writer.Write("Hello world")
}

func main() {
	pl := fmt.Println
	f, err := os.Create("test.txt")
	if err != nil {
		pl("Failed to create file")
		return
	}
	defer f.Close()

	WriteHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
