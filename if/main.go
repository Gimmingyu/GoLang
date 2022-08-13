package main

import (
	"fmt"
	"math/rand"
)

func HasRichFriend() bool {
	a := rand.Int()
	if a > 10 {
		return true
	}
	return false
}

func GetFriendsCount() int {
	return rand.Int()
}

func main() {
	const value int = 10
	const deny string = "신발끈이 풀렸네"

	// value = 20 => Compile error
	// var value int = 20 ==> Compile error

	price := 35000

	if price >= 50000 {
		if HasRichFriend() {
			fmt.Println(deny)
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 {
		if GetFriendsCount() > 3 {
			fmt.Println(deny)
		} else {
			fmt.Println("나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다")
	}

	// if 초기문, 조건문 예시
	// if filename, success := UploadFile(); success {
	// 	fmt.Println("Upload success", filename)
	// } else {
	// 	fmt.Println("Failed to upload")
	// }
}
