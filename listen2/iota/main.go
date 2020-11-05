package main

import "fmt"

func main() {
	const (
		aa = iota
		bb
		cc
	)

	const (
		a1 = 1 << iota
		b1
		c1
	)

	fmt.Println("aa=%d bb=%d cc=%d", aa, bb, cc)
	fmt.Println("a1=%d b1=%d c1=%d", a1, b1, c1)

}
