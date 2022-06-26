package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "干了兄弟萌"
	//fmt.Printf("%d ", s)
	//fmt.Println(utf8.RuneCountInString(s))
	//bytes := []byte(s)
	//for len(bytes) > 0 {
	//	_, size := utf8.DecodeRune(bytes)
	//	fmt.Println("index: ", size)
	//	bytes = bytes[size:]
	//	fmt.Println(bytes)
	//}
	//for i, ch := range []rune(s) {
	//	fmt.Println(i, ch)
	//}
	//for i, ch := range []byte(s) {
	//	fmt.Println(i, ch)
	//}
	res := []rune(s)
	fmt.Println(string(res[1:]))
	fmt.Println(strings.Trim("  s   ss         1 ", " "))
}
