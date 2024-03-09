package main

import (
	"fmt"
	"math"
)

type Point4 struct {
	x, y float64
}

func (p *Point4) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point4
	name string
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

type Customer2 struct {
	Name string
	Log
}

type Customer3 struct {
	Name string
	log  Log
}

func main() {
	n := &NamedPoint{Point4{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())

	c := new(Customer)
	c.Name = "Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"

	c = &Customer{"Obama", &Log{"1 - Yes we can!"}}
	fmt.Println(c.log)
	// c.log.Add("2 - after me")
	c.Log().Add("2 - after me")
	fmt.Println(c.Log())
	fmt.Println()

	c2 := &Customer2{"Obama", Log{"5 - Yes we can"}}
	c2.Add("6 - aboba")
	fmt.Println(c2)

	c3 := &Customer3{"Obama", Log{"10 - yes"}}
	c3.log.Add("11 - yes")
	fmt.Println(c3)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.Log)
}

func (c *Customer3) String() string {
	return c.Name + "\nLog:\n" + fmt.Sprintln(c.log)
}
