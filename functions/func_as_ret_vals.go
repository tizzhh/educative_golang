package main

import "fmt"

func main() {
	p2 := Add2()
	fmt.Println(p2(3))
	twoadder := Adder(2)
	fmt.Println(twoadder(5))
	f := Adder2()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
