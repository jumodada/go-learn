package main

import (
	"fmt"
	"time"
)

func worker() chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println("received", <-c)
		}
	}()
	return c
}
func channelDemo() {
	for i := 1; i < 10; i++ {
		c := worker()
		c <- i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}
