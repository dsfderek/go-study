package main

import "fmt"

func testIf1() {
	num := 11
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func testIf2() {
	num := 21
	if num >= 3 && num < 10 {
		fmt.Printf("num:%d num >3 && num <10\n", num)
	} else if num >= 10 && num < 20 {
		fmt.Printf("num:%d num >10 && num <20\n", num)
	} else if num >= 20 && num < 30 {
		fmt.Printf("num:%d num >20 && num <30\n", num)
	}
}

func testIf3() {

	if num := 11; num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	//fmt.Printf("num:%d\n",num)
}

func testIf4() {

	if num := getNum(); num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	//fmt.Printf("num:%d\n",num)
}

func getNum() int {
	return 13
}

func main() {
	//testIf1()
	//testIf2()
	//testIf3()
	testIf4()
}
