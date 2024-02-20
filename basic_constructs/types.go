package main

import "fmt"

func main() {
	var a int
	var b int32
	a = int(b)
	fmt.Printf("%T: %05d %T\n", a, a, b)
	var aboba float64 = 5.5
	fmt.Printf("%g\n", aboba)
	var c complex64 = 5 + 10i
	fmt.Printf("%T %v\n", c, c)
	var ch byte = 'A'
	fmt.Printf("%c\n", ch)
}
