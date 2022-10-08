package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello", i)
			}

		}(i)
	}
	time.Sleep(time.Millisecond)
}
