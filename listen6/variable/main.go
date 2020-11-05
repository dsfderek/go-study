package main

import "fmt"

// 全局变量
var a int = 100

// 局部变量 --> 函数内部有效
func test() {
	b := 200
	fmt.Println(b)
}

func main() {
	test()
}
