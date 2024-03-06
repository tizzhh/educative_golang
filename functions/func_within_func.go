package main

import "fmt"

func f1(a, b, c int) int {
	return a + b + c
}

func f2(a, b int) (int, int, int) {
	n1 := a + b
	n2 := a - b
	n3 := a * b
	return n1, n2, n3
}

func main() {
	fmt.Println(f1(f2(20, 10)))
	aboba := f1
	fmt.Println(aboba(1, 2, 3))
}
