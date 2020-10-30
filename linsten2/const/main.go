package main

import "fmt"

func main() {
	/*const a int = 123
	const b string = "hello"*/
	const (
		a int    = 123
		b string = "hello"
		c int    = 200
		d
		e int = 300
		f
	)

	fmt.Println("a=%d\n", a)
	fmt.Println("b=%s\n", b)
	fmt.Println("c=%d\n", c)
	fmt.Println("d=%d\n", d)
	fmt.Println("e=%d\n", e)
	fmt.Println("f=%d\n", f)
}
