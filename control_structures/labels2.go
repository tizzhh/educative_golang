package main

import "fmt"

func main() {
	// i := 0
	// HERE:
	// fmt.Printf("%d\n", i)
	// i++
	// if i == 5 {
	// 	return
	// }
	// goto HERE

LABEL1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				goto LABEL1 // jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
