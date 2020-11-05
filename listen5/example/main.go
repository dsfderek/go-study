package main

import "fmt"

// 求1到100之内的所有质数，打印
func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
func example1() {
	for i := 2; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("[%d] 是 质数\n", i)
		}
	}
}

func is_shuixianhuashu(n int) bool {
	// 个位
	first := n % 10
	// 十位
	second := (n / 10) % 10
	// 百位
	third := (n / 100) % 10

	sum := first*first*first + second*second*second + third*third*third

	//fmt.Printf("n=%d first=%d second=%d third=%d\n", n, first, second, third)

	//return false
	if n != sum {
		return false
	}
	return true
}

func example2() {
	for i := 100; i < 1000; i++ {
		if is_shuixianhuashu(i) == true {
			fmt.Printf("[%d] 是 水仙花数\n", i)
		}
	}
}

// 输入一行字符，分别统计出其中英文字母、空格、数字和其他字符的个数
func cacl_num(str string) (charCount int, numCount int, spaceCount int, otherCount int) {

	// 转成utf8的字符
	uftChars := []rune(str)
	for i := 0; i < len(uftChars); i++ {
		// 是不是英文字符
		if uftChars[i] >= 'a' && uftChars[i] <= 'z' || uftChars[i] >= 'A' && uftChars[i] <= 'Z' {
			charCount++
			continue
		}
		// 是不是数字
		if uftChars[i] >= '0' && uftChars[i] <= '9' {
			numCount++
			continue
		}

		// 是不是空格
		if uftChars[i] == ' ' {
			spaceCount++
			continue
		}

		otherCount++

	}
	return
}

func example3() {
	str := "12345  fdfe^%$#@@@@@222222"
	charCount, numCount, spaceCount, otherCount := cacl_num(str)
	fmt.Printf("charCount=%d numCount=%d spaceCount=%d otherCount=%d\n", charCount, numCount, spaceCount, otherCount)

}

func main() {
	//example1()
	//example2()
	example3()
}
