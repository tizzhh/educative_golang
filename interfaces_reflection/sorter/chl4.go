package main

import (
	"fmt"
	"sorter/mysort"
)

type Person struct {
	firstName, lastName string
}

type Persons []Person

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
	return p[i].firstName+p[i].lastName < p[j].firstName+p[j].lastName
}

func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	data := []Person{Person{"z", "z"}, Person{"x", "x"}, Person{"aboba", "valeriy"}}
	persons := Persons(data)
	fmt.Println(persons)
	sorter.Sort(persons)
	fmt.Println(persons)
}
