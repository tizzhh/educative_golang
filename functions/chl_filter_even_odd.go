package main

import "fmt"

type flt func(int) bool

func isEven(n int) bool {
	return n%2 == 0
}

func filter(sl []int, f flt) (yes, no []int) {
	for _, val := range sl {
		if f(val) {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return yes, no
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filter(sl, isEven))
}
