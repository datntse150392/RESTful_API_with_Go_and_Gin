package main

import "fmt"

func main() {

	// Khai báo một mảng primes gồm 6 phần tử kiểu interger
	primes := [6]int{2, 3, 5, 7, 11} 

	// Khai báo một slice s gồm 2 phần tử là từ index = 1 tới index 4-1
	var s []int = primes[1:4]
	fmt.Println(s)

	// Chú ý khi khai báo slide thì nó sẽ không tạo ra một ô nhớ mới, và nó chỉ trỏ tới ô nhớ của array mà nó cắt đi
	// Nếu chúng ta có thay đổi một element của slice thì bên mảng bị slice cắt đi cũng sẽ bị thay đổi theo
	// Fx:
	names:= [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	// Khai báo một slice a từ index = 0 tới index = 2
	a := names[0:2]
	// Khai báo một slice b từ index = 1 tới index = 3
	b := names[1:3]
	fmt.Println(a, b)
	// Thay đổi element thứ 1 của slice b
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)


	// Slice literals
	// Slice literals giống như array literals nhưng không cần phải khai báo độ dài của slice
	// Slice literals được viết bằng cách liệt kê các giá trị của slice, được bao quanh bởi dấu ngoặc nhọn {}.
	// Fx:
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)


	// Slice defaults
	// Khi slice được khai báo mà không có index thì mặc định index = 0 và index = len(array) - 1
	// Fx:
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	n = n[:2]
	fmt.Println(n)
	n = n[1:]
	fmt.Println(n)
}