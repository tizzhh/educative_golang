package main

import (
	"fmt"
	"math"
)

type E struct {
	thing int
}

func (e *E) change() {
	e.thing = 1
}

func (e E) write() string {
	return fmt.Sprint(e)
}

type Point3 struct{ x, y, z float64 }

func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) { *l = append(*l, val) }

type T4 struct {
	a int
}

func (t T4) print(message string) {
	fmt.Println(message, t.a)
}

func (T4) hello(message string) {
	fmt.Println("Hello!", message)
}

func callMethod(t T4, method func(T4, string)) {
	method(t, "A message")
}

func main() {
	var e1 E
	e1.change()
	fmt.Println(e1.write())

	e2 := new(E)
	e2.change()
	fmt.Println(e2.write())

	p3 := &Point3{3, 4, 5}
	fmt.Println(p3.Abs())

	var lst List
	lst.Append(1)
	fmt.Println(lst, lst.Len())

	plst := new(List)
	plst.Append(2)
	fmt.Println(plst, plst.Len())

	t1 := T4{10}
	t2 := T4{20}
	var f func(T4, string) = T4.print
	callMethod(t1, f)
	callMethod(t2, f)
	callMethod(t1, T4.hello)
}
