package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("main开始运行")
	ch := make(chan int, 1)

	for i := 1; i <= 5; i++ {
		go work(i, ch)
	}

	ret := <-ch
	log.Printf("第%d个用户最快完成任务", ret)
}

func work(no int, ch chan int) {
	i := rand.Intn(10) + 1
	fmt.Printf("用户%d完成任务需要%d秒钟\n", no, i)
	time.Sleep(time.Duration(i) * time.Second)
	ch <- no
}
