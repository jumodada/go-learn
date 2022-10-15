package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1", n)
		case n := <-c2:
			fmt.Println("Received from c2", n)
		}
	}
}
