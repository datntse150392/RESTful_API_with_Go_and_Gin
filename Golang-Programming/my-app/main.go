package main

import "fmt"

func main() {
	var i interface{}

	i = 42
	fmt.Println(i)


	i = "hello"
	fmt.Println(i)
}
