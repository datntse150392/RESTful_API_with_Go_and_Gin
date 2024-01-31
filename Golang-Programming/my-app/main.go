package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	println(hypot(5, 2))

	fmt.Println(compute(math.Pow))

	pos, reg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			reg(-2*i),
		)
	}

	// => Function closure được hiểu là trong một block code, có một function được khai báo trên trong một function khác thì function được khai báo sẽ có thể tham chiếu biến biến ở bên ngoài function nhưng trong phạm vi block code đó thôi
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 2)
}

func adder() func(int) int {
	sum := 0;
	return func(x int) int {
		sum += x
		return sum
	}
}