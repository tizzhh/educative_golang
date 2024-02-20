package main

import "fmt"

func main() {
	var test = `пвтопaboba\naboba`
	fmt.Println(test)
	fmt.Println(len(test))
	s1 := "aboba"
	s2 := "privet"
	s3 := s1 + s2
	fmt.Println(s3)
}
