package main

import "fmt"

type Animal struct {
	name string
}

func (animal Animal) Sing() {
	fmt.Println(animal.name + " Sing a Song.")
}

func (animal Animal) Jump() {
	fmt.Println(animal.name + " Jump the river.")
}
