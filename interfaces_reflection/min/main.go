package main

import (
	"fmt"
	"miner/min"
)

func ints() {
	a := min.IntArray([]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586})
	m := min.Min(a)
	fmt.Println("min:", m)
}

func strings() {
	a := min.StringArray([]string{"ddd", "eee", "bbb", "ccc", "aaa"})
	fmt.Println("min string:", min.Min(a))
}

func main() {
	ints()
	strings()
}
