package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("type is %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("type is %T\n", u)
	} else {
		fmt.Println("area does not contain contain a var of type Circle")
	}
}