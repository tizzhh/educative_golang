package main

import "fmt"

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
		"lima": 98}
)

func main() {
	invMap := make(map[int][]string, len(barVal))
	for k, v := range barVal {
		// if _, isPresent := invMap[v]; isPresent {
		// 	invMap[v] = append(invMap[v], ",")
		// }
		invMap[v] = append(invMap[v], k)
	}

	fmt.Println(invMap)

}
