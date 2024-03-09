package main

import "fmt"

type Square struct {
	side float32
}

type Triangle struct {
	base, height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface {
	Perimiter() float32
}

func main() {
	intf := AreaInterface(&Square{69})
	fmt.Println(intf.Area())
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimiter() float32 {
	return 4 * sq.side
}

func (tr *Triangle) Area() float32 {
	return 0.5 * tr.base * tr.height
}
