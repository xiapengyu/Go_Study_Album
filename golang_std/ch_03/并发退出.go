package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool)  {
	for  {
		select {
		default:
			fmt.Println("正常工作")
		case <- ch:
			fmt.Println("退出工作")
			return
		}
	}
}

func main()  {
	cancel := make(chan bool)

	for i := 0; i < 10; i++ {
		go worker(cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
}
