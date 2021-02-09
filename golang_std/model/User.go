package main

import "fmt"

type User struct {
	name string
	age int
	address string
}

func(user User) Sing() {
	fmt.Println(user.name + " Sing a Song.")
}

func (user User) Jump()  {
	fmt.Println(user.name + " Jump the river.")
}
