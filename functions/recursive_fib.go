package main

import "fmt"

func main() {
	res := 0
	for i := 0; i <= 10; i++ {
		res = fibonacci(i)
		fmt.Printf("fib(%d) is %d\n", i, res)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
