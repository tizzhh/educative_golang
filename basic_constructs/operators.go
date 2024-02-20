package main

import "fmt"

func main() {
	var a int = 5
	a++
	fmt.Println(a)
	b3 := 10 > 5 // greater than operator
	fmt.Println(b3)
	b3 = 10 < 5 // less than operator
	fmt.Println(b3)
	b3 = 5 <= 10 // less than equal to
	fmt.Println(b3)
	b3 = 10 != 10 // not equal to
	fmt.Println(b3)
	var b int = 6
	fmt.Println(a ^ b)
}
