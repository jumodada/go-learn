package main

import "fmt"

func main() {
	m1 := map[string]int{
		"ss": 1,
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
