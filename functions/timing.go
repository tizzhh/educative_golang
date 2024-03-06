package main

import (
	"fmt"
	"time"
)

func Calc() {
	for i := 0; i < 10000; i++ {

	}
}

func main() {
	start := time.Now()
	Calc()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
