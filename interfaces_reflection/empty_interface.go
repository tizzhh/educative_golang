package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool: // if v is bool
			fmt.Printf("any %v is a bool type", v)
		case int: // if v is int
			fmt.Printf("any %v is an int type", v)
		case float32: // if v is float32
			fmt.Printf("any %v is a float32 type", v)
		case string: // if v is string
			fmt.Printf("any %v is a string type", v)
		case specialString: // if v is specialString
			fmt.Printf("any %v is a special String!", v)
		default: // none of types satisfied
			fmt.Println("unknown type!")
		}
	}
	testFunc(whatIsThis)
}

type Element interface{}

type Vector struct {
	a []Element
}

func (p *Vector) At(i int) Element {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}

func main() {
	var val Any
	val = 5
	fmt.Println("val value:", val)
	val = str
	fmt.Println("val value:", val)

	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Println("val value:", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("type in %T\n", t)
	case string:
		fmt.Printf("type string %T\n", t)
	case *Person:
		fmt.Printf("type person %T\n", t)
	default:
		fmt.Println("unknown")
	}

	TypeSwitch()

	fmt.Println()
	v := Vector{}
	v.a = append(v.a, "string")
	v.a = append(v.a, 5.5)
	fmt.Println(v.At(0), v.At(1))
}
