package main

import (
	"fmt"
)

func source(ch chan<- int) {
	ch <- 1
}

func sink(ch <-chan int) {
	fmt.Println(<-ch)
}

func main() {
	var c = make(chan int)
	go source(c)
	go sink(c)
}
