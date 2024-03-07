package main

import "fmt"

var fib [10]int64

func fibs() [10]int64 {
	start, prev := int64(0), int64(1)
	for i := 0; i < len(fib); i++ {
		fib[i] = prev
		prev, start = prev+start, prev
	}
	return fib
}

func main() {
	fmt.Println(fibs())
}
