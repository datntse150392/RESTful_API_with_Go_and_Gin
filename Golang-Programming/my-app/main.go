package main

func main() {
	a := make([]int, 5)
	println("a", a)

	b := make([]int, 0, 10)
	println("b", b)

	c := b[0:2]
	println("c", c)

	d := c[2:5]
	println("d", d)

	e := d[2:5]
	println("e", e)

}