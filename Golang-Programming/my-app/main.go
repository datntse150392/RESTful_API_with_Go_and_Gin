package main

import (
	"fmt"
)
func main() {
	fmt.Println("Hello, World!")
	
	// Khai báo channel unbuffered (không có buffer) 
	// ch := make(chan int)

	// Khai báo channel buffered (có buffer)
	// ch2 := make(chan int, 2) // Có thể chứa tối đa 2 phần tử
	// ch2 <- 1
	// ch2 <- 2
	// ch2 <- 3 // Lỗi, vì đã đầy deadlock

	// Ví dụ về channel buffer
	ch3 := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- i
			fmt.Print("Send: ", i, "\n")
		}
		close(ch3)
	}()

	for v := range ch3 {
		fmt.Println("Receive: ", v)
	}

	/*
		Giải thích close and range
		close:
		- Khi một channel được đóng, không thể gửi dữ liệu vào nó nữa
		- Khi một channel được đóng, vẫn có thể nhận dữ liệu từ nó
		- Khi một channel được đóng, không thể đóng nó lần nữa
		- Khi một channel được đóng, không thể mở nó lên

		range:
		- range trên channel sẽ lặp qua tất cả các phần tử trong channel cho đến khi channel đóng
		- Khi channel đóng, range sẽ dừng lặp
	*/

}
