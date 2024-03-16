package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go goro(ch)
	fmt.Println("sending")
	ch <- 10
	fmt.Println("sent")
}

func goro(ch chan int) {
	fmt.Println("aboba")
	time.Sleep(5 * time.Second)
	fmt.Println("recieved", <-ch)
}
