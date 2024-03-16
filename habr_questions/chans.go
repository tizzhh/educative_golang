package main

import (
	"fmt"
	"time"
)

func evenSum(from, to int, ch chan int) {
	res := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			res += i
		}
	}
	ch <- res
}

func squareSum(from, to int, ch chan int) {
	res := 0
	for i := 0; i <= to; i++ {
		if i%2 == 0 {
			res += i * i
		}
	}
	ch <- res
}

func evenSumsync(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			res += i
		}
	}
	return res
}

func squareSumsync(from, to int) int {
	res := 0
	for i := 0; i <= to; i++ {
		if i%2 == 0 {
			res += i * i
		}
	}
	return res
}

func main() {
	// start := time.Now()
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 10000000, evenCh)
	go squareSum(0, 10000000, sqCh)
	// select {
	// case <- evenCh:
	// 	fmt.Println("x")
	// case y := <- sqCh:
	// 	fmt.Println("y", y)
	// }
	// fmt.Println(<-evenCh + <-sqCh)
	// fmt.Println(time.Now().Sub(start))

	// start = time.Now()
	// res1 := evenSumsync(0, 1000000000)
	// res2 := squareSumsync(0, 1000000000)
	// fmt.Println(res1 + res2)
	// fmt.Println(time.Now().Sub(start))

	for {
		select {
		case x := <- evenCh:
			fmt.Println(x)
			return

		case y := <- sqCh:
			fmt.Println(y)
			return
		default:
			fmt.Println("not accessible")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
