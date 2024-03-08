package main

import "fmt"

func main() {
	mapLit := map[string]int{"one": 1, "two": 2}
	var mapAssigned map[string]int
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	// mapAssigned := mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Println(mapLit["one"], mapCreated["key2"], mapAssigned["two"], mapLit["ten"])
	fmt.Println(mapLit, mapAssigned, len(mapLit))

	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)

	map2 := make(map[string]float32, 100)
	fmt.Println(len(map2))

	mp1 := make(map[int][]int)
	mp2 := make(map[int]*[]int)

	mp1[1] = append(mp1[1], 3)
	mp1[1] = append(mp1[1], 4)

	// *mp2[1] = append(*mp2[1], 10)
	// *mp2[1] = append(*mp2[1], 11)

	fmt.Println(mp1, mp2)

}
