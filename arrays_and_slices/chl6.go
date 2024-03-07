package main

import "fmt"

func Filter(s []int, f func(int) bool) (res []int) {
	for _, val := range s {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func even(x int) bool {
	return x%2 == 0
}

func main() {
	fmt.Println(Filter([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, even))
}
