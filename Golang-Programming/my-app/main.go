package main

import "math"

type Vertex struct {
	X, Y int
}

// Method
func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X)*float64(v.X)+float64(v.Y)*float64(v.Y))
}

func main() {
	v := Vertex{1,2}
	println(v.Abs())
}