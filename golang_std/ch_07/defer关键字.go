package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ") //参数已经传入，按照后入先出顺序执行，输出1 2 3
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ") //参数未传入，按照后入先出顺序执行，输出0 0 0
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(num int) {
			fmt.Print(num, " ") //参数已经传入，按照后入先出顺序执行，输出1 2 3
		}(i)
	}
}

func main() {
	d1()
	fmt.Println()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
