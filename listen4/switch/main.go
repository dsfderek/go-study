package main

import "fmt"

func testIf() {
	var a int = 4

	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("a=5\n")
	}
}
func testSwitch1() {
	//var a int  = 4
	a := 2
	switch a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	}
}

func testSwitch2() {
	//var a int  = 4
	switch a := 10; a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	default:
		fmt.Printf("错误\n")
	}
}

func getValue() int {
	return 8
}

func testSwitch3() {
	switch a := getValue(); a {
	case 1, 2, 3, 4, 5:
		fmt.Printf("a>=1 and a <= 5\n")
	case 6, 7, 8, 9, 10:
		fmt.Printf("a>=6 and a <= 10\n")
	default:
		fmt.Printf("a>11\n")
	}
}

func testSwitch4() {
	num := 102
	switch {
	case num >= 0 && num < 60:
		fmt.Printf("num=%d 不及格\n", num)
	case num >= 60 && num < 70:
		fmt.Printf("num=%d 及格\n", num)
	case num >= 70 && num < 80:
		fmt.Printf("num=%d 良好\n", num)
	case num >= 80 && num < 90:
		fmt.Printf("num=%d 优秀\n", num)
	case num >= 90 && num < 100:
		fmt.Printf("num=%d 卓越\n", num)
	case num == 100:
		fmt.Printf("num=%d 满分\n", num)
	default:
		fmt.Printf("num=%d 分数错误\n", num)
	}
}

func testMulti() {
	// 1 * 1 = 1
	// 1 * 2 = 2  2 * 2 = 4
	// 1 * 3 = 3  2 * 3 = 6  3 * 3 = 9
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", j, i, i*j)
		}
		fmt.Println()

	}
}

func main() {
	//testIf()
	//testSwitch1()
	//testSwitch2()
	//testSwitch3()
	//testSwitch4()
	testMulti()
}
