package main

import (
	"fmt"
)

func swap(x , y string) (string, string) {
	return y, x
}




func main() {
	a,b := swap("hello", "world")
	fmt.Println(a,b)

	// Khai báo biến có 2 cách
	var i,j int = 1,2
	// Biến python sẽ tự hiểu type của biến
	python := "short"
	fmt.Println(i,j,python)
}