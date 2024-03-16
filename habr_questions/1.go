package main

import (
	"fmt"
	"time"
)

func main() {
	// var x int
	// fmt.Scan(&x)
	// switch x {
	// case 1:
	// 	fmt.Println("aboba")
	// case 2, 3:
	// 	fmt.Println("aboba2")
	// }

	// aboba := new(map[string]string)
	// *aboba = map[string]string{}
	// (*aboba)["aboba"] = "privet"
	// fmt.Println((*aboba)["aboba"])

	// var inter interface{} = 1
	// fmt.Printf("%T %v %v\n", inter, inter, inter == 1)

	// inter = 2
	// var (
	// 	// v string
	// 	ok bool
	// )
	// _, ok = inter.(int)
	// fmt.Println(ok)

	// const (
	// 	UNKNOWN = iota
	// 	FEMALE
	// 	MALE
	// )
	// fmt.Println(UNKNOWN)

	// nums := 1 << 5
	// defer fmt.Println(nums)
	// nums = nums >> 1
	// fmt.Println("done")

	ch := make(chan bool)

	go out(0, 5, ch)
	go out(6, 10, ch)

	<-ch
	<-ch
}

func out(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- true
}
