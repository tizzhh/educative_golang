package main

import "fmt"

type flt func(int) bool

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

// Same comment for isEven
func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func filter(sl []int, f flt) []int {
	var res []int
	for _, val := range sl {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}
