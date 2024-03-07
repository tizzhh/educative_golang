package main

import "fmt"

func sum(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}
	return s
}

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))

	slice1 := make([]int, 10)
	fmt.Println(slice1, cap(slice1))
	slice1 = append(slice1, 1)
	fmt.Println(slice1, cap(slice1))
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

	var slice3 []int = make([]int, 10)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = 5 * i
	}
	fmt.Println(slice3, len(slice3), cap(slice3))

	var p *[]int = new([]int)
	fmt.Println(*p)
}
