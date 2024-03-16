package main

import "fmt"

func download(to int, ch chan int) {
	res := 0
	for i := 0; i <= to; i++ {
		res += i
	}
	ch <- res
}

func main() {
	ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)
	fmt.Println(<-ch1 + <-ch2 + <-ch3)
}