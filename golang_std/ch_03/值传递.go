package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)

	func() {
		x[0] = 7
		fmt.Println(x)
	}()

	var code string
	for i := 0; i < 6; i++ {
		n := rand.Intn(10)
		code += strconv.Itoa(n)
	}
	fmt.Println(code)

	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	//code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	//fmt.Println(code)

	//var array = []int{1, 2, 3, 4, 5, 6}
	//var code strings.Builder
	//for _, v := range array {
	//	fmt.Println(v)
	//	code.WriteString(strconv.Itoa(v))
	//}
	//fmt.Println(code.String())

	//num := 19
	//a := string(num)
	//b := strconv.Itoa(num)
	//fmt.Printf("%T, %T", a, b)

	//注意string()与strconv.Itoa()的区别
	h := 98
	fmt.Println(string(h))
}
