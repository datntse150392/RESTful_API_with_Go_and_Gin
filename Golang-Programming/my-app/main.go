package main

import (
	"fmt"
	"time"
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
		for i := 0; i < 100; i++ {
			ch3 <- i
			fmt.Print("Send: ", i, "\n")
		}
	}()

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch3)
		time.Sleep(time.Second)
	}

	/*
	Giải thích ví dụ trên
		+ Do khi khai báo channel ch3 đã có buffer 10, tức là có thể nhận 10 slot giá trị trước khi bị block
		+ Khi gửi giá trị vào channel ch3, nếu channel chưa đầy thì sẽ không bị block
		+ Khi bắt đầu vòng lặp thứ 11 thì phải đợi lấy giá trị từ channel ch3, nếu channel chưa có giá trị thì sẽ bị block
	*/
}
