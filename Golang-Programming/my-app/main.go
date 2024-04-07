package main

import (
	"fmt"
)
func main() {
	fmt.Println("Hello, World!")
	// ch := make (chan int)

	// VD1:
	// fmt.Print("VD1:", <-ch)

	// VD2: 
	ch2 := make(chan int)

	go func () {
		ch2 <- 1
	}()

	// ch2 <- 2
	fmt.Print("VD2:", <-ch2)

}
