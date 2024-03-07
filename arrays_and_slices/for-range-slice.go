package main

import (
	"fmt"
)

// import "strings"

func main() {
	// var slice1 []int = []int{1, 2, 3, 4}
	// for ix, val := range slice1 {
	// 	fmt.Printf("%d %d\n", ix, val)
	// }

	// seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	// for _, season := range seasons {
	// 	fmt.Println(season)
	// }

	// for ix := range seasons {
	// 	seasons[ix] = strings.ToUpper(seasons[ix])
	// }
	// fmt.Println(seasons)

	// val := 0
	// screen := [2][2]int{}

	// for row := range screen {
	// 	for col := range screen[row] {
	// 		screen[row][col] = val
	// 		val++
	// 	}
	// }

	// for row := range screen {
	// 	for column := range screen[0] {
	// 		fmt.Print(screen[row][column], " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	ar := [...]int{1, 2, 3}
	for _, val := range ar {
		val *= 2
	}
	fmt.Println(ar)

	sl := []int{1, 2, 3}
	fmt.Println(len(sl[1:1]))
}
