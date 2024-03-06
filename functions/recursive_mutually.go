package main

import "fmt"

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
}

func even(n int) bool {
	if n == 0 {
		return true
	}
	return odd(AbsVal(n) - 1)
}

func odd(n int) bool {
	if n == 0 {
		return false
	}
	return even(AbsVal(n) - 1)
}

func AbsVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
