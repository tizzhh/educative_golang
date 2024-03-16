package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)

	// ch1 := make(chan int)
	// go pump(ch1)
	// go suck(ch1)
	// fmt.Println(<-ch1)
	time.Sleep(1e9)
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
	close(ch)
}
