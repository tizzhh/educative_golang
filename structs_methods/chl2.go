package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Abs())
}
