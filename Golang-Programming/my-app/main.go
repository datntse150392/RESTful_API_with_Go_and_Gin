package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)

	n := 10000
	queueCh := make(chan int, n)
	maxWorkers := 5

	for i := 1; i <= n; i++ {
		queueCh <- i
	}

	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go func() {
			for v := range queueCh {
				time.Sleep(20 * time.Microsecond)
				fmt.Println("Processing job", v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}