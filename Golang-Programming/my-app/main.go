package main

import "fmt"

func main() {
	var a [2]string // Biến a là mảng string có 2 phần tử
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Print(a)
}