package main

import "fmt"

func arraySlice() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[1:])
	fmt.Println(arr[:4])
	fmt.Println(arr[1:3])
	fmt.Println(arr[:])
}

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	//var arr1 [5]int
	//arr2 := [3]int{1, 2, 3}
	//var grid [3][4]int
	//arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(arr2, arr1, arr3)
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	//for _, v := range arr3 {
	//	fmt.Println(v)
	//}
	//var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	////updateSlice(a[1:2])
	//s1 := a[3:6]
	//s2 := s1[0:6]
	//fmt.Println(s1)
	//fmt.Println(s2)

	//var s []int
	//for i := 0; i < 100; i++ {
	//	s = append(s, i)
	//}
	//fmt.Println(s)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{9, 8, 7}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	slice2 = append(slice2, slice1...)
	fmt.Println(slice1, slice2)
	fmt.Println(slice2[len(slice2)-1])
}
