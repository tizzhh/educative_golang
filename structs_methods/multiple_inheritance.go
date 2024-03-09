package main

import (
	"fmt"
	"strconv"
)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type SmartPhone struct {
	Camera
	Phone
}

type Integer int

func (i *Integer) String() string {
	return strconv.Itoa(int(*i))
}

func main() {
	cp := new(SmartPhone)
	fmt.Println(cp.TakeAPicture(), cp.Call())

	var a Integer
	fmt.Println(a)
}
