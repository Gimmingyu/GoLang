package main

import (
	"fmt"
)

type User struct {
	id       int
	name     string
	password string
}

type UserInterface interface {
	UpdateUser(password string) User
	GetUser()
}

func AddUser(id int, name string, password string) User {
	newUser := User{
		id:       id,
		name:     name,
		password: password,
	}
	return newUser
}

func (user User) UpdateUser(password string) User {
	newUser := User{
		id:       user.id,
		name:     user.name,
		password: password,
	}
	return newUser
}

func (user User) GetUser() {
	fmt.Println(user.name, user.id)
}

func main() {
	pl := fmt.Println
	u1 := User{
		id:       1,
		name:     "John",
		password: "password",
	}
	u2 := AddUser(u1.id+1, "mingkim", "1234")
	var ui UserInterface
	ui = u1
	ui.GetUser()
	pl("u1 ", u1)
	pl("u2 ", u2)
	u1.UpdateUser("5678")
	u2.UpdateUser("9999")
	pl("u1 ", u1)
	pl("u2 ", u2)
}
