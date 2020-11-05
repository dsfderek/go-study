package main

import "fmt"

func testReverseStringV1() {
	str := "helloe"
	bytes := []byte(str)
	fmt.Println("len ", len(str))
	fmt.Println("bytes ", bytes)

	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		fmt.Println("tmp:", tmp)
		fmt.Println("bytes[i]", bytes[i])
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp

	}

	fmt.Println("bytes ", bytes)
	str2 := string(bytes)

	fmt.Println(str2)
}

func testReverseStringV2() {
	str := "hello中文"
	r := []rune(str)
	fmt.Println("len ", len(str))
	fmt.Println("r ", r, len(r))

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		fmt.Println("tmp:", tmp)
		fmt.Println("r[i]", r[i])
		r[len(r)-i-1] = r[i]
		r[i] = tmp

	}
	//
	//fmt.Println("bytes ", bytes)
	str2 := string(r)
	//
	fmt.Println(str2)
}

func testHuiWen() {
	str := "上海自来水来自海上"
	r := []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	//
	//fmt.Println("bytes ", bytes)
	str2 := string(r)
	if str == str2 {
		fmt.Println("is ok")
	} else {
		fmt.Println("is not ok")
	}

}

func main() {
	//testReverseStringV1()
	//testReverseStringV2()
	testHuiWen()
}
