package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	v := math.Pow(x,n)
	if v < lim {
		return v
	} else {
		return 0
	}
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(pow(3,2,10))

	switch os:= runtime.GOOS; os {
	case "darwin": fmt.Println("macOS")
	case "linux": fmt.Println("Linux")
	default: 
		fmt.Println(os)
	}

	t:= time.Now()
	switch{
		case t.Hour() < 12: fmt.Println("Good Morning")
		case t.Hour() < 12: fmt.Println("Good Afternoon")
		default: fmt.Println("Good Evening")
	}
}