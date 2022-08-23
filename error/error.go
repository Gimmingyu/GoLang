package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("Failed to create file ", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file ", err)
			return
		}
	}
	fmt.Println("File description :", line)
}

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}
