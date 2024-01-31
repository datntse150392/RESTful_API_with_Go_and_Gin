package main

import "fmt"

type Vertex struct {
	Lat, Log float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	n := make(map[string]int)
	n["Answer"] = 42
	fmt.Println("The value:", n["Answer"])

	n["Answer"] = 48
	fmt.Println("The value:", n["Answer"])

	delete(n, "Answer")
	fmt.Println("The value:", n["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// => khi delete thì sẽ xóa key đó luôn
}
