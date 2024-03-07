package main

import "fmt"
import "sort"

func main() {
	// c := []byte("aboba")
	// fmt.Println(c)

	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	fmt.Println()

	r := []rune("привет")
	fmt.Println(r, len(r), "привет", string(r))

	s1 := "hello"
	c := []byte(s1)
	c[0] = 'c'
	s1 = string(c)
	fmt.Println(s1)

	isl := []int{9, 8, 7}
	sort.Ints(isl)
	fmt.Println(isl, sort.IntsAreSorted(isl))

	isl2 := []int{1, 2, 3, 4}
	isl2 = append(isl2[:2], append([]int{69}, isl2[2:]...)...)

	var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

	arru := make([]byte, len(arr))
	ixu := 0
	tmp := byte(0)
	for _, val := range arr {
		if val != tmp {
			arru[ixu] = val
			fmt.Printf("%c ", arru[ixu])
			ixu++
		}
		tmp = val
	}
	fmt.Println(string(arru))
}
