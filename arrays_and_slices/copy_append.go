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
}
