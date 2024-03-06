package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	callback(1, Add)
	s := "Hello! Let's run Go lang"
	fmt.Println(strings.IndexFunc(s, unicode.IsSpace))
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
