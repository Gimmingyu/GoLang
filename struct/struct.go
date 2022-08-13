package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User
	VIPLevel int
	Price    int
}

func main() {

	pl := fmt.Println

	var a Student = Student{
		Name:  "Mingkim",
		Class: 0,
		No:    0,
		Score: 0.0,
	}

	pl(a)

	var house House
	house.Address = "성남시 수정구"
	house.Size = 28
	house.Price = 28
	house.Category = "Apartment"
	pl(house)
	house.Price = 50
	pl(house)

	user := User{"Mingyu Kim", "mingkim", 27}
	vip := VIPUser{user, 3, 250}
	pl(user, vip)
}
