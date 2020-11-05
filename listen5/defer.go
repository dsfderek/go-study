package main

import "fmt"

func testDefer1(){
	defer fmt.Println("hello")
	defer fmt.Println("hello v1")
	defer fmt.Println("hello v2")
	defer fmt.Println("hello v3")
	fmt.Println("aaaa")
	fmt.Println("bbbb")
}

func testDefer2(){
	defer fmt.Println("start")
	for i:=0;i<5;i++ {
		defer fmt.Printf("i=%d\n",i)
	}
	defer fmt.Println("需求")
	defer fmt.Println("end")
}


func testDefer3(){
	i := 0
	defer fmt.Printf("defer i=%d\n",i)
	i = 1000
	fmt.Printf("i==%d\n",i)
}

func main(){
	//testDefer1()
	//testDefer2()
	testDefer3()
}

