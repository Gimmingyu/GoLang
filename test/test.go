package main

import (
	"fmt"
	"log"
)

type CustomError struct {
	Message string
	Code    string
}

func logFatal(err error) {
	defer RecoverSuccess()
	log.Panicf("My custom error, a is %v\n", err)
}

func logPanic(err error) {
	defer RecoverSuccess()
	log.Panicf("My custom error, a is %v\n", err)
}

func RecoverSuccess() {
	if r := recover(); r != nil {
		log.Printf("Recover Success... Error was %v\n", r)
	}
}

func Last(err error) {
	defer log.Println("Last Function 지나가요~")
	log.Panicf("My custom error, a is %v\n", err)
}

func BringDeepAndDeeper(err error) {
	defer log.Println("BringDeepAndDeeper 지나가요~")
	Last(err)
}

func BringDeeper() {
	msg := fmt.Errorf("BringDeeper made this error")
	defer log.Println("BringDeeper 지나가요~")
	BringDeepAndDeeper(msg)
}

func main() {

	logFatal(fmt.Errorf("Fatal Error"))
	logPanic(fmt.Errorf("Panic Error"))
	fmt.Println("here!!!!")

	fmt.Println("New test!!!!!!!!!!")

	defer RecoverSuccess()
	BringDeeper()
}
