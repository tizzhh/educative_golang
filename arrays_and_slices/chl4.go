package main

import "fmt"

func enlarge(s []int, factor int) []int {
	res := make([]int, len(s)*factor)
	copy(res, s)
	return res
}

func main() {
	fmt.Println(enlarge([]int{1, 2, 3}, 5))
}
