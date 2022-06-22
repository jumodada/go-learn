package main

import "fmt"

func variable() {
	var a, b int = 1, 2
	var s string = "Hello World"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	a, b, s := 1, 2, "Hello World"
	fmt.Println(a, s, b)
}

func main() {
	variable()
	variableTypeDeduction()
}
