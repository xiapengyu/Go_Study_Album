package main

import "fmt"

/**
返回自然数序列的管道
*/
func GenerateNatural() chan int {
	var in = make(chan int)
	go func() {
		for i := 2; ; i++ {
			in <- i
		}
	}()
	return in
}

func PrimeFilter(in chan int, prime int) chan int {
	var out = make(chan  int)
	go func() {
		for  {
			if i :=<- in; i%prime !=0{
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}
