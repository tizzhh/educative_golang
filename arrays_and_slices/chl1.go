package main

import "fmt"

func main() {
	var arrray [15]int
	for i := 0; i < len(arrray); i++ {
		arrray[i] = i
	}
	fmt.Println(arrray)
}
