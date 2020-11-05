package main

import (
	"fmt"
	"github.com/gostudy/linsten2/access"
	"strings"
)



func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	a = !a

	fmt.Println(a)

	var b bool = true

	if a == true && b == true {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	if a == true || b == true {
		fmt.Println("|| ok")
	} else {
		fmt.Println("|| not ok")
	}

	fmt.Println("b=%t", b)
}

func testInt() {
	var a int8 // 127 ~ -127
	a = 18
	fmt.Println("a = ", a)
	a = -12
	fmt.Println("a = ", a)
	//a = 127
	//fmt.Println("a = ", a)

	var b int
	b = 4343224234242342
	fmt.Println("b=", b)

	//var c  float32  = 5.677
	var c = 5.677
	fmt.Println("c=", c)

	//b = int(a)
	fmt.Println("a=%d  a=%x c=%f\n", a, a, c)
}

func testString() {
	//var a string = "dsf"
	var a = "hello a"
	fmt.Println(a)

	b := a
	fmt.Println(b)

	c := "hello c"
	fmt.Println(c)

	//c := "x"
	fmt.Println("a=%v b=%v c=%v\n", a, b, c)

	d := `
	床前明月光，
	疑是地上霜
	`
	//fmt.Println("d=%v\n", d)
	//
	//cLen := len(d)
	//
	//fmt.Println("len cLen=%s", cLen)

	d = d + d

	//fmt.Println("d=%s", d)
	// 拼接
	d = fmt.Sprintf("%s%s", d, d)
	fmt.Println("d=%s\n", d)

	// 切割
	ips := "10.108.23.22;10.108.23.24"
	ipSlice := strings.Split(ips, ";")
	fmt.Println("first in :%s\n", ipSlice[0], "\n", ipSlice[1])

	// 包含
	result := strings.Contains(ips, "10.108.23.22")

	fmt.Println("result=%s", result)

	// 前缀和后缀
	str := "http://baidu.baidu.com"

	if strings.HasPrefix(str, "http") {
		fmt.Println("str is http url")
	} else {
		fmt.Println("str is not  http url")
	}

	if strings.HasSuffix(str, ".com") {
		fmt.Println("str is baidu url")
	} else {
		fmt.Println("str is not baidu url")
	}

	// 子串第一次出现的位置
	index := strings.Index(str, "baidu")
	fmt.Println("baidu is index:%d\n", index)
	index = strings.LastIndex(str, "baidu")
	fmt.Println("baidu is index:%d\n", index)

	// join
	var strArr []string = []string{"10.123.9.1", "10.123.9.3", "10.123.9.4"}
	resultArr := strings.Join(strArr, ";")

	fmt.Println("resultArr= %s\n", resultArr)

}

func testOperator(){
	var a int = 2
	if a == 2 {
		fmt.Println("is ok\n")
	}else{
		fmt.Println("is ont ok")
	}

}

func testAccess(){
	fmt.Println("access.a=%d\n", access.A)
}
func main() {
	//testBool()
	//testInt()
	//testString()
	//testOperator()
	testAccess()
}
