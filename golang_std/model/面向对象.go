package main

func main() {
	user := User{name: "Tom", age: 12, address: "China"}
	user.Jump()
	user.Sing()

	Sing(User{name: "Lucy", age: 12, address: "China"})
	Sing(Animal{name: "JoJo"})
}

func Sing(shape Shape) {
	shape.Jump()
	shape.Sing()
}
