package main

import "fmt"

func main() {
	il := 5
	fmt.Printf("%d: %p\n", il, &il)
	var intP *int = &il
	fmt.Println(intP, *intP)
	fmt.Println(il == *(&il))

	*intP = 10
	fmt.Println(intP, *intP, il)

	var p *int = nil // Making a nil pointer
	fmt.Println(p)
}
