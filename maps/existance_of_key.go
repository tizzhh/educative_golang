package main

import "fmt"

func main() {
	mp := map[string]int{"hello": 1}
	if val, ok := mp["hello"]; ok {
		fmt.Println(val)
	}

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25

	value, isPresent = map1["Beijing"]

	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println("does not exist")
	}

	value, isPresent = map1["Paris"]
	fmt.Println(value, isPresent)

	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println("does not exist")
	}
}
