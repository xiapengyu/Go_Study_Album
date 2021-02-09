package main

import (
	"fmt"
	"sync"
	"time"
)

var RWLock sync.RWMutex
var h int = 10

/**
读写锁，多个线程可以同时读，但是写的时候其他线程既不可以写也不可以读
*/
func main() {
	//for i := 0; i < 10; i++ {
	//	go read(i)
	//}
	write(1)
	readF(2)
	write(3)
}

func read(i int) {
	fmt.Printf("%d read start \n", i)
	RWLock.RLock()
	fmt.Printf("%d reading %d \n", i, h)
	RWLock.RUnlock()
	fmt.Printf("%d read complete \n", i)
}

func readF(i int) {
	println(i, "read start")
	RWLock.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	RWLock.RUnlock()
	println(i, "read over")
}

func write(i int) {
	println(i, "write start")
	RWLock.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	RWLock.Unlock()
	println(i, "write over")
}
