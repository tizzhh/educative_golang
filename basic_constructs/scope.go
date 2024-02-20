package main

import "fmt"

var num = 5

func main() {
	var dec = true
	fmt.Printf("%d\n", num)
	num = 10
	fmt.Printf("%v\n", num)
	fmt.Printf("\t%t\n", dec)
}
