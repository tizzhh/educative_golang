package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chf    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chf {
		f()
	}
}

func (p *Person) SetSalary(sal float64) {
	p.chf <- func() { p.salary = sal }
}

func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chf <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	bs := NewPerson("Smith bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("salary changes:")
	fmt.Println(bs)
}
