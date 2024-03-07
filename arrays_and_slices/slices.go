package main

import "fmt"

func main() {
	// s := []int{1, 2, 3}
	// fmt.Println(s)
	// a := s
	// a[0] = 10
	// fmt.Println(s)
	// fmt.Println(a)
	// var arr1[7]int
	// var slice1 []int = arr1[2:5]
	// fmt.Println(slice1, arr1)
	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i
	// }
	// fmt.Println(slice1, arr1)
	// fmt.Println(len(arr1), len(slice1), cap(slice1))
	// slice1 = slice1[0:4]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1), cap(slice1))

	// var slice []int
	// fmt.Println(slice, len(slice), cap(slice))
	// slice = append(slice, 1)
	// fmt.Println(slice, len(slice), cap(slice))
	// slice = append(slice, 2)
	// fmt.Println(slice, len(slice), cap(slice))
	// slice = append(slice, 3)
	// fmt.Println(slice, len(slice), cap(slice))
	// slice = append(slice, 4)
	// fmt.Println(slice, len(slice), cap(slice))
	// slice = append(slice, 5)
	// fmt.Println(slice, len(slice), cap(slice))

	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)
	var arr = [2]int{1, 2}
	var slice []int = arr[:]
	fmt.Println(slice, arr)
	slice = append(slice, 3)
	fmt.Println(slice, arr)
}
