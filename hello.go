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

//func convertToBinary(n int) string {
//	result := ""
//	for ; n > 0; n /= 2 {
//		lsb := n % 2
//		result = strconv.Itoa(lsb) + result
//	}
//	return result
//}
//
//func printFile(fileName string) {
//	file, err := os.Open(fileName)
//	if err != nil {
//		panic(err)
//	}
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		fmt.Println(scanner.Text())
//	}
//}

//func function(a, b int) (q, r int) {
//	q = a / b
//	r = a % b
//	return q, r
//}

//func apply(cb func(a, b int) int, a, b int) int {
//	return cb(a, b)
//}
//
//func closures() func(num int) int {
//	var state = 1
//	return func(num int) int {
//		state += num + 1
//		return state
//	}
//}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	//variable()
	//forTest()
	//fmt.Println(function(10, 2))
	//fmt.Println(apply(func(a, b int) int {
	//	return a + b
	//}, 3, 2))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	//var cb = closures()
	//fmt.Println(cb(2))
	//fmt.Println(cb(2))
	//fmt.Println(convertToBinary(10))
}
