package main

import (
	"fmt"
	"go-learn/start/queue"
)

func main() {
	q := queue.Queue{1}
	x := queue.Queue{2}
	x.Push(33)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
}
