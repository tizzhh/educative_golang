package main

import "fmt"

func fibarray(term int) []int {
	res := make([]int, term)

	prev, start := 0, 1
	for i := 0; i < term; i++ {
		res[i] = prev
		prev, start = start, prev+start
	}
	return res
}

func main() {
	fmt.Println(fibarray(5))
}
