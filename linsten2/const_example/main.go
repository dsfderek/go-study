package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d
		e = 300
		f
		g = iota
		h
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h, "\n")

	const (
		aa = iota
		bb
	)

	fmt.Println(aa)
	fmt.Println(bb)
}
