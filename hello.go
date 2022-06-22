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

func constTest() {
	const a = 1
	fmt.Println(a)
	{
		const a = 2
		fmt.Println(a)
	}
}

func main() {
	//variable()
	constTest()
}
