package main

import "fmt"

func bubbleSort(sl []int) {
	for i := 1; i < len(sl); i++ {
		for j := 0; j < len(sl)-i; j++ {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}
}

func main() {
	ab := []int{4, 5, 2, 1, 3}
	bubbleSort(ab)
	fmt.Println(ab)
}
