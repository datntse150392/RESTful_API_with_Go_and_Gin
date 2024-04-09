package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0

	lock := new(sync.Mutex)

	for i := 1; i <= 10; i++ {

		go func() {
			for j := 1; j <= 1000; j++ {
				lock.Lock()
				count++
				fmt.Println(count)
				lock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 5)
}