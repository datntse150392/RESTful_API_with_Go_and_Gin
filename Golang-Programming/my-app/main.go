package main

import "fmt"

func main() {
var a *int
b := 10
a = &b

c:= 2701
fmt.Println("a:", a) // In ra giá trị ô nhớ của biến a
fmt.Println("a:", *a) // In ra giá trị của biến b thông qua con trỏ a
a = &c;
fmt.Println("a:", *a) // In ra giá trị ô nhớ của biến a

}