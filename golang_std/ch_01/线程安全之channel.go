package main

import (
	"fmt"
	"sync"
)

var m int = 0
var group sync.WaitGroup

func main() {
	c := make(chan int, 1)
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go addInt(c)
	}
	group.Wait()
	fmt.Println(m)
}

func addInt(c chan int) {
	defer group.Done()
	c <- 1
	m++
	<-c
}
