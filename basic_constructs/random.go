package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Int()
	b := rand.Intn(8)
	fmt.Printf("%d %d\n", a, b)
}
