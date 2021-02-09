package main

import "fmt"

func main() {
	//golang中没有while，循环都用for来实现，go不愿意引入太多关键字，为了语言简洁

	//一般的for循环
	fmt.Println("一般的for循环")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	//实现while循环
	fmt.Println("实现while循环")
	i := 0
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}

	//实现do...while循环
	fmt.Println("现do...while循环")
	n := 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if n > 10 {
			anExpression = false
		}
		fmt.Print(n, " ")
		n++
	}

}
