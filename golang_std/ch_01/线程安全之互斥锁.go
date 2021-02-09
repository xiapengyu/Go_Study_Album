package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

//使用sync.Mutex互斥锁实现线程安全
var mutex sync.Mutex
var n int = 0

func main() {
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go addNum()
	}

	w.Wait()
	fmt.Println(n)
}

func addNum() {
	mutex.Lock()
	defer mutex.Unlock()
	defer w.Done()
	n++
}
