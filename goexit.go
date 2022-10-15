package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("test1 end")
	runtime.Goexit()
	fmt.Println("test1 begin")
}
func main() {
	go func() {
		fmt.Println("test begin")
		test()
		fmt.Println("test end")
	}()
}
