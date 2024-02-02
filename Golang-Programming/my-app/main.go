package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y int
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() int {
	return int(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func Scale(v *Vertex, f int) {
	v.X = v.X * f
	v.Y = v.Y * f
} 

type Rectangle struct {
	width, height float64
}

// Phương thức thay đổi trạng thái của đối tượng với receiver tham trị
func (r Rectangle) ScaleValue(factor float64) {
	r.width *= factor
	r.height *= factor
}

// Phương thức thay đổi trạng thái của đối tượng với receiver con trỏ
func (r *Rectangle) ScalePointer(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	// Writing code follow the method
	v := Vertex{3,5};
	v.Scale(10);
	println(v.X, v.Y);
	println(v.Abs());


	// Writing code follow the function
	Scale(&v, 10)
	println(v.X, v.Y);
	println(v.Abs());


	// Pointer indirection
	p := &v
	println(p.Abs())
	
	r1 := Rectangle{3, 4}

  // Sử dụng phương thức với receiver tham trị
  r1.ScaleValue(2)
  fmt.Println("Trạng thái sau khi ScaleValue:", r1) // Kết quả: {3 4}, không thay đổi

  // Sử dụng phương thức với receiver con trỏ
  r1.ScalePointer(2)
  fmt.Println("Trạng thái sau khi ScalePointer:", r1) // Kết quả: {6 8}, đã thay đổi
}