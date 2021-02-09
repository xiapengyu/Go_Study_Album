package main

import (
	"fmt"
)

func main() {
	Array()
}

//具名函数
func Array() {
	var a = []int{1, 2, 3} // a 是一个数组
	var b = &a             // b 是指向数组的指针

	for i := range a {		//for range方式迭代的性能可能会更好一些，不用担心下标越界
		fmt.Println(i)
	}

	for i, v := range a {
		fmt.Printf("Array[%d]=%d\n", i, v)
	}

	for i, v := range *b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	//匿名函数
	var add = func(a, b int) int {
		return a + b
	}

	fmt.Println(add(4, 6))

	//defer语句延迟执行，通常是在行数执行完之后执行，
	//每个defer语句延迟执行的函数引用的都是同一个i迭代变量，
	//在迭代体中谨慎使用defer语句
	for i := 0; i < 3; i++ {
		defer func(){ println(i) } ()
	}

	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func(){ println(i) } ()
	}

	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int){ println(i) } (i)
	}
}


