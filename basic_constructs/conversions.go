package main

import "fmt"

func main() {
	var number = 5.2
	fmt.Println(number)
	fmt.Println((int(number)))
	const (
		UNKNOWN = iota
		FEMALE
		MALE
	)
	fmt.Println(FEMALE)
}
