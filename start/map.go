package main

import "fmt"

func main() {
	m1 := map[string]string{
		"luna": "moon",
	}
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}
	delete(m1, "luna")
	val, hasLion := m1["lion"]
	fmt.Println(val, hasLion)
	fmt.Println(m1["luna"])
}
