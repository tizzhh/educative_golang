package main

import "fmt"

func insertSlice(slice, insertion []string, index int) []string {
	// var res []string
	// res = append(res, slice[:index]...)
	// res = append(res, insertion...)
	// res = append(res, slice[index:]...)
	// return res

	res := make([]string, len(slice)+len(insertion))
	ix := copy(res, slice[:index])
	ix += copy(res[ix:], insertion)
	copy(res[ix:], slice[index:])
	return res
}

func main() {
	fmt.Println(insertSlice([]string{"M", "N", "O", "P", "Q", "R"}, []string{"A", "B", "C"}, 0))
}
