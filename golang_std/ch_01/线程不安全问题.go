package main

import (
	"fmt"
	"sync"
)

var num int = 0
var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()
	fmt.Println(num)
}

func add() {
	num++
	defer wg.Done()
}
