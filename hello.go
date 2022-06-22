package main

import "fmt"

//var c = 222

//var (
//	xx int = 1
//	yy int = 1
//	zz     = 1
//)

//func variable() {
//	var a, b int = 1, 2
//	var s string = "Hello World"
//	fmt.Println(a, s, b)
//}
//
//func variableTypeDeduction() {
//	a, b := 1, 2
//	//a = b + 2
//	fmt.Println(a, b, c)
//	fmt.Println(xx, yy, zz)
//}

//func constTest() {
//	const a = 1
//	fmt.Println(math.Sqrt(64))
//	{
//		const a = 2
//		fmt.Println(a)
//	}
//}

//func constTest() {
//	const (
//		FILENAME = 1
//		b        = 2
//	)
//	fmt.Println(math.Sqrt(64))
//	{
//		const a = 2
//		fmt.Println(a)
//	}
//}

func enum() {
	//const (
	//	lenovo    = 0
	//	bytedance = 1
	//	baidu     = 2
	//	alibaba   = 3
	//	tencent   = 4
	//)

	const (
		js = iota*10 + 1
		_
		golang
		java
	)
	fmt.Println(js, golang, java)
}

func main() {
	//variable()
	enum()
}
