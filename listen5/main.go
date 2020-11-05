package main

import "fmt"

func sayHello() {
	fmt.Println("sey hello world")
}

//func add(a int, b int) int {
//	return a + b
//}

func add2(a, b, c int) int {
	return a + b + c
}

func Computer(X, Y int) (int, int) {
	a := X + Y
	b := X * Y
	return a, b
}

//func ddd(a, b int) (int, int) {
//	sum := a + b
//	sub := a - b
//	return sum, sub
//}

//func calc(a, b int) (int, int) {
//	sum := a + b
//	sub := a - b
//	return sum, sub
//}

func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func cacl_v1(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum
}

func calc_v2(a int, b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum += b[i]
	}
	return sum + a
}

func main() {
	//sayHello()
	//s := add(3,4)

	//s2 := add2(2, 3, 4)
	//
	//fmt.Println(s2)

	//d, f := Computer(100, 200)
	//fmt.Println(d, f)

	//sum, sub := ddd(200, 100)
	//fmt.Println(sum, sub)

	//sum, _ := calc(200,100)
	//calc(200,100)
	//fmt.Println()

	//sum := cacl_v1(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//fmt.Println(sum)

	sum := calc_v2(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum)

}
