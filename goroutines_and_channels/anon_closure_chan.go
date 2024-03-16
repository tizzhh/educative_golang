package main

import "fmt"

func prod(c chan func()) {
	f := <-c
	f()
	c <- nil
}

func main() {
	c := make(chan func())
	go prod(c)
	x := 0
	c <- func() {
		x++
	}
	fmt.Println(x)
}
