package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// The slice grows on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2,3,4,5,6)
	printSlice(s)

	s = append(s, 8,9,0)
	printSlice(s)

	// => Tóm lại là hàm append() sẽ tạo ra một slice mới và copy các phần tử từ slice cũ sang slice mới
}

func printSlice(s []int) {
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}