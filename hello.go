package main

import "fmt"

var c = 222

var (
	xx = 1
	yy = 1
	zz = 1
)

func variable() {
	var a, b int = 1, 2
	var s string = "Hello World"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	a, b := 1, 2
	//a = b + 2
	fmt.Println(a, b, c)
	fmt.Println(xx, yy, zz)
}

func main() {
	//variable()
	variableTypeDeduction()
}
