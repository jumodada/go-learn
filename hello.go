package main

import "strconv"

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

//func enum() {
//	//const (
//	//	lenovo    = 0
//	//	bytedance = 1
//	//	baidu     = 2
//	//	alibaba   = 3
//	//	tencent   = 4
//	//)
//
//	const (
//		js = iota*10 + 1
//		_
//		golang
//		java
//	)
//	fmt.Println(js, golang, java)
//}

//func ifStatement() {
//	const fileName = "ab.txt"
//	if _, err := ioutil.ReadFile(fileName); err != nil {
//		fmt.Println("get failed")
//	}
//	else {
//		fmt.Println("got it")
//	}
//}

//func switchTest() int {
//	var i, result = rand.Intn(100), 1
//	fmt.Println(i)
//	switch {
//	case i > 90:
//		result = 1
//	case i > 80:
//		result = 2
//	case i > 60:
//		result = 3
//	case i > 50:
//		result = 4
//	}
//	return result
//}

//func forTest() int {
//	sum := 1
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	return sum
//}

func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	//variable()
	//forTest()
	convertToBinary(10)
}
