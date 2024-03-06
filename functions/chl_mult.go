package main

import "fmt"

func main() {
	fmt.Println(SumProductDiff(3, 4))
}

func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}
