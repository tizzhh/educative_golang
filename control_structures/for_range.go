package main

import "fmt"

func main() {
	str := "go is cool!"

	for pos, char := range str {
		fmt.Printf("char on pos %d is %c ", pos, char)
	}
	fmt.Println()

	str2 := "Japanese: 日本語"
	for pos, char := range str2 {
		fmt.Printf("char on pos %d is %c ", pos, char)
	}
	fmt.Println()
}
