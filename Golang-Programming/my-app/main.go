package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var(
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{} // X:0 and Y:0
	p = &Vertex{1, 2} // has type *Vertex
)

func main() {

	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)

	// Using piointer
	p:= &v // Dùng con trỏ p trỏ đến ô nhớ địa chỉ của v
	p.X = 10 // Thay đổi giá trị của v thông qua con trỏ p
	fmt.Println(v.X)

	z := &Vertex{1, 2} // has type *Vertex
	fmt.Println(z.X)
}