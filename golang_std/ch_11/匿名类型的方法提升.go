package main

import (
	"fmt"
)

type People struct {
	Age    int
	gender string
	Name   string
}

type OtherPeople struct {
	People
}

func (p People) PeInfo() {
	fmt.Println("People ", p.Name, ": ", p.Age, "岁, 性别:", p.gender)
}

func main() {
	//特别注意方法提升的规则
	other := OtherPeople{People{10, "S", "H"}}
	other.People.PeInfo()
	other.PeInfo()

	cal()
}

func cal() {
	t := 100000
	a := 78350
	b := 3000
	fmt.Println(t - a - b)

	//垃圾回收与SetFinalizer
	//var m runtime.MemStats
	//runtime.ReadMemStats(&m)
	//log.Printf("%d Kb\n", m.Alloc/1024)

	fmt.Printf("%.2f \n", 1.2558)
	sprint := fmt.Sprintf("%.2f \n", 123.358)
	fmt.Println(sprint)

}
