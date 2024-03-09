package main

import "fmt"

type Shaper interface {
	Area() float32
	// Permiter() float32 square does not implement that method, illegal
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	lenght, width float32
}

func (r *Rectangle) Area() float32 {
	return r.lenght * r.width
}

// other methods can be implemented
func (r *Rectangle) aboba() {
	fmt.Println("hello")
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Println(areaIntf.Area())

	r := &Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	for _, fig := range shapes {
		fmt.Println(fig, fig.Area())
	}
}
