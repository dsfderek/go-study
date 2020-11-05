package main

import "fmt"

func main() {
	var (
		a int
		b bool
		c string
		d float32
	)
	fmt.Println("a=%d b=%t c=%s d=%f\n", a, b, c, d)

	a = 10
	b = true
	c = "hello"
	d = 10.8
	fmt.Println("a=%d b=%t c=%s d=%f\n", a, b, c, d)
}
