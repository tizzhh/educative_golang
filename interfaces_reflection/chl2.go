package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	x int
}

func (s *Simple) Get() int {
	return s.x
}

func (s *Simple) Set(n int) {
	s.x = n
}

type RSimple struct {
	x int
}

func (p *RSimple) Get() int {
	return p.x
}

func (p *RSimple) Set(n int) {
	p.x = n
}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(69)
	case *RSimple:
		it.Set(420)
	}
	return it.Get()
}

func main() {
	inf := Simpler(&RSimple{1})
	res := fI(inf)

	fmt.Println(res)
}
