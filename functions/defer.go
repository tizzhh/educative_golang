package main

import "fmt"

func main() {
	Function1()
}

func Function1() {
	fmt.Println("func1 top")
	defer Function2()
	fmt.Println("func1 bottom")
}

func Function2() {
	fmt.Println("func2")
}
