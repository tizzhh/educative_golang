package main

import "fmt"

func main() {
	// slice1 := make([]int, 0, 10)
	// fmt.Println(slice1)
	// for i := 0; i < cap(slice1); i++ {
	// 	slice1 = slice1[0:i + 1]
	// 	slice1[i] = i
	// 	fmt.Println(slice1)
	// 	fmt.Printf("len is %d\n", len(slice1))
	// }
	// fmt.Println(slice1)

	ar := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	fmt.Println(a)
	a = a[0:4]
	fmt.Println(a)

	// test := make([]int, 0, 50)
	// fmt.Println(test[25])
}
