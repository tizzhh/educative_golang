package main

import (
	"fmt"
	"strings"
)

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := fib()
	for i := 0; i <= 10; i++ {
		fmt.Println(i+2, f())
	}

	addBmp := MakeAddSuffix(".bmp")
	// addJpeg := MakeAddSuffix(".jpeg")

	fmt.Println(addBmp("file"))

}
