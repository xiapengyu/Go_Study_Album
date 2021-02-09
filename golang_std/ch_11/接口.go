package main

import "fmt"

type B interface {
	fn()
}

type A struct {
	name string
}

func (a A) fn() {
	fmt.Println(a.name)
}

type C int

func (c C) fn() {
	fmt.Println("c.fn", c)
}

func main() {
	var a, b A = A{name: "LeBron James"}, A{name: "Kobe Bryant"}
	a.fn()
	b.fn()

	var v B = A{name: "Tracy MacGrady"} // 接口类型可接受结构体A的值，因为结构体A实现了接口
	v.fn()


	var m, n C = 10, 99
	m.fn()
	n.fn()

	var o B = C(67)
	o.fn()

}
