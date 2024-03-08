package main

import "fmt"

type innserS struct {
	in1, in2 int
}

type outerS struct {
	b int
	c float32

	int
	innserS
}

type A struct{ a int }
type B struct{ a, b int }
type C struct {
	A
	B
} //ambiguity

type D struct {
	B
	b float32
}

type K1 struct {
	a int
}

type K2 struct {
	b int
}

type K3 struct {
	a int
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Println(outer)

	outer2 := outerS{6, 7.5, 60, innserS{5, 10}}
	fmt.Println(outer2)

	var c C
	fmt.Println(c.A.a, c.B.a)

	d := D{B{5, 5}, 5.5}
	fmt.Println(d.b, d.B.b)

	var person struct {
		name, surname string
	}
	person.name, person.surname = "Barack", "obama"

	anotherPerson := struct {
		name, surname string
	}{"Barack", "Obama"}
	fmt.Println(person, anotherPerson)

	t1 := K1{10}
	t2 := K2{10}
	t3 := K3{10}
	t4 := K3{20}
	fmt.Println(t1, t2, t3, t4)
	// fmt.Println("t1 == K2?", t1==t2) // <-- invalid operation: t1 == K2 (mismatched types K1 and K2)
	// fmt.Println("t1 == K3?", t1==t3) // <-- invalid operation: t1 == K3 (mismatched types K1 and K3)
	fmt.Println("t1 == t3?", t1 == K1(t3))
	fmt.Println("K3 == t4?", t3 == t4)
}
