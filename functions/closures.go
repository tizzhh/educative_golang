package main

import "fmt"

func main() {
	fmt.Println(func(x, y int) int { return x + y }(5, 10))

	// f()

	fv := func(s string) {
		fmt.Println("hey ", s)
	}
	fv("aboba")
	fmt.Printf("type of fv is %T\n", fv)
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Println(i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
