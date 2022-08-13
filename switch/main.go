package main

import (
	"fmt"
	"math/rand"
)

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func main() {

	var a int32 = int32(rand.Int())
	p := fmt.Println

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1, 2, 3")
	}

	b := 1

	switch true {
	case b < 10:
		fmt.Println("Smaller than 10")
		fallthrough
	case b < 100:
		fmt.Println("Smaller than 100")
		fallthrough
	case b < 1000:
		fmt.Println("Smaller than 1000")
		fallthrough
	case b < 10000:
		fmt.Println("Smaller than 10000")
		fallthrough
	case b < 1000000:
		fmt.Println("Smaller than 1000000")
		fallthrough
	case b < 100000000:
		fmt.Println("Smaller than 100000000")
		fallthrough
	case b < 1000000000:
		fmt.Println("Smaller than 1000000000")
		fallthrough
	default:
		fmt.Println("Big Enough")
	}

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) age는 현재 스코프에서 없는 값이기 때문에 에러가 발생한다.

	p("My favorite color is", colorToString((getMyFavoriteColor())))

}

func getMyAge() int {
	return 22
}

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}
