package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker_(ctx, &wg)
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()

}

func worker_(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return
		}
	}

}
