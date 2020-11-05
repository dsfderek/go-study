package calc

var a int = 100 // 私有
var A int = 200 // 可以被外部的包访问

func Add(a,b int) int{
	return a +b
}

func sub(a,b int) int{
	return  a-b
}