package main

import "fmt"
import "unicode"

func main() {
	var ch1 rune = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 rune = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch1, ch2, ch3) // UTF-8 code point
	fmt.Println(unicode.IsDigit(ch1))
}
