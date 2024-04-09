package main

import (
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}

}

func main() {

	// Fx 1
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func (){
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()

	go func (){
		time.Sleep(1 * time.Second)
		ch2 <-2
	}()

	// fmt.Print("ch1", <-ch1)
	// fmt.Print("ch2", <-ch2)	
	select{
	case v1 := <-ch1: 
		println("ch1", v1)
	case v2 := <-ch2:
		println("ch2", v2)
	}

	// Fx2
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
