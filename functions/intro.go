package main

import "fmt"

func main() {
	fmt.Println("main")
	greeting("privet")
	fmt.Println("after main")
}

func greeting(name string) {
	fmt.Println("hellooo", name)
	name = "aboba"
	fmt.Println("hellooo", name)
}
