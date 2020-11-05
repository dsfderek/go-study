package main

import (
	"fmt"
	"github.com/gostudy/listen6/calc"
)

func main(){
	s1 := 100
	s2 := 200
	sum := calc.Add(s1,s2)
	fmt.Printf("sum=%d\n",sum)
}
