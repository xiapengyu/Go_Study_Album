package main

import "fmt"

type I interface {
	fn()
	Mn()
}

type T string

func (t T) fn() {
	fmt.Println(t)
}

//方法接收器是方法的拥有者与调用者，可以是类型或者是类型的指针，但不能是指针类型
func (t T) Mn(){

}

// invalid receiver type Q (Q is a pointer type)
//type MyInt int
//type Q *MyInt
//func (q Q) print() {
//	fmt.Println("Q:", q)
//}





type Stringer interface {
	String() string
}

func main() {
	var varI I = T("String")
	if v, ok := varI.(T); ok {
		fmt.Println("varI类型断言结果为：", v) // varI已经转为T类型
		varI.fn()
	}

	var value interface{}
	switch str := value.(type) {
	case string:
		fmt.Println("value类型断言结果为string:", str)
	case Stringer:
		fmt.Println("value类型断言结果为Stringer:", str)
	default:
		fmt.Println("value类型不在上述类型之中")
	}

	value = "类型断言检查"
	str, ok := value.(string)
	if ok {
		fmt.Printf("value类型断言结果为：%T\n", str) // str已经转为string类型
	} else {
		fmt.Printf("value不是string类型 \n")
	}

}
