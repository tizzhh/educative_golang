package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to, n)
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	sl3 = append(sl3, sl_to...)
	fmt.Println(sl3)

	a := [5]int{0, 2, 4, 6, 8}
	// var s []int = a[1:3]
	// fmt.Println(s, a)
	// s = append(s, 10, 10, 10, 10, 10)
	// fmt.Println(s, a)


	slice := a[1:3:3]
	fmt.Println(slice, a)
	slice = append(slice, 10)
	fmt.Println(slice, a)
	// s = append(s, 10)
	// s = append(s, 10)
	// fmt.Println(s, a)
}
