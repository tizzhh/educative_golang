package main

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	// for i := 0; i < HEIGHT; i++ {
	// 	for j := 0; j < WIDTH; j++ {
	// 		screen[j][i] = pixel(i * j)
	// 	}
	// }

	// fmt.Println(screen)

	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[0][0])
	fmt.Println(values)

}
