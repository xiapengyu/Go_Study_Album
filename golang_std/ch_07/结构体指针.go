package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func NewPerson(name string, age int, address string) Person {
	return Person{name, age, address}
}

func NewPersonPointer(name string, age int, address string) *Person {
	return 	&Person{name, age, address}
}

func main() {
	p1 := NewPerson("tom", 15, "HuBei")
	fmt.Println(p1.name)
	fmt.Println(p1)

	p2 := NewPersonPointer("tom", 15, "HuBei")
	fmt.Println(p2.name)
	fmt.Println(p2)
	fmt.Println(*p2)

	//const r1 = '€'
	const r1 = 'a'
	fmt.Println("(int32) r1:", r1)
	fmt.Printf("(HEX) r1: %x\n", r1)
	fmt.Printf("(as a String) r1: %s\n", r1)
	fmt.Printf("(as a character) r1: %c\n", r1)
}
