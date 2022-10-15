package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello", i)
		}
	}()

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("Gosched")
	}
}
