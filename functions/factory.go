package main

import "fmt"

type flt func(int) bool
type slice_split func([]int) ([]int, []int)

func isOdd(num int) bool {
	return num%2 != 0
}

func isBiggerThan4(num int) bool {
	return num > 4
}

func filter_factory(f flt) slice_split {
	return func(s []int) (yes, no []int) {
		for _, val := range s {
			if f(val) {
				yes = append(yes, val)
			} else {
				no = append(no, val)
			}
		}
		return yes, no
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)
	odd_even_function := filter_factory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd = ", odd)
	fmt.Println("even = ", even)
	//separate those that are bigger than 4 and those that are not.
	bigger, smaller := filter_factory(isBiggerThan4)(s)
	fmt.Println("Bigger than 4: ", bigger)
	fmt.Println("Smaller than or equal to 4: ", smaller)
}
