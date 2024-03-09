package main

import "fmt"

type AnonStruct struct {
	a float32
	int
	string
}

func main() {
	aS := AnonStruct{1, 2, "3"}
	fmt.Println(aS)
}
