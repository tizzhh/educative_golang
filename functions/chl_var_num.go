package main

import "fmt"

func main() {
	fmt.Println(sumInts(5, -2, 0, 9))
}

func sumInts(list ...int) (sum int) {
	for _, val := range list {
		sum += val
	}
	return sum
}
