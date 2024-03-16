package main

import (
	"fmt"
	"time"
)

func main() {
	// stream := pump()
	// go suck(stream)
	producer_consumer()
	time.Sleep(15 * time.Microsecond)
}

func producer_consumer() {
	suck(pump())
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	// go func() {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}
	// }()
}
