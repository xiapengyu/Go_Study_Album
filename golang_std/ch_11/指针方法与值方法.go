package main

import "fmt"

type S struct {
	name string
}

func (s S) M1() {
	s.name = "name1"
}

func (s *S) M2() {
	s.name = "name2"
}

func main() {
	str := S{"Kobe Bryant"}
	fmt.Printf("调用M1前的值%s \n", str.name)
	str.M1()
	fmt.Printf("调用M1后的值%s \n", str.name)

	fmt.Printf("调用M2前的值%s \n", str.name)
	str.M2()
	fmt.Printf("调用M2后的值%s \n", str.name)

	t2 := &S{"t2"}
	fmt.Println("M1调用前：", t2.name)
	t2.M1()
	fmt.Println("M1调用后：", t2.name)

	fmt.Println("M2调用前：", t2.name)
	t2.M2()
	fmt.Println("M2调用后：", t2.name)

	//调用M1前的值Kobe Bryant
	//调用M1后的值Kobe Bryant
	//调用M2前的值Kobe Bryant
	//调用M2后的值name2
	//M1调用前： t2
	//M1调用后： t2
	//M2调用前： t2
	//M2调用后： name2

	//方法：值传递传递的是值的一份拷贝，无法修改原来的值，地址传递才可以修改变量的值
	//如果要在方法内部修改变量的值，需要定义指针方法，使用地址传递

}