package main

import "fmt"

type Simpler interface {
	Get() int
	Set(n int)
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

func fI(it Simpler) int {
	it.Set(69)
	return it.Get()
}

func main() {
	inf := Simpler(&Simple{10})
	fmt.Println(fI(inf))

	var aboba float32 = -3.0
	fmt.Println(aboba / 0)

}
