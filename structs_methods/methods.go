package main

import (
	"fmt"
	"time"
)

type S struct {
	a int
}

type SType S
type SAlias = S
type IntType int
type IntAlias = int

func (recv S) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}

func (recv SType) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv SAlias) print() {
// 	fmt.Printf("%t: %[1]v\n", recv)
// }

func (recv IntType) print() {
	fmt.Printf("%t: %[1]v\n", recv)
}

// func (recv IntAlias) print() {
// 	fmt.Printf("%t: %[1]v\n", recv)
// }

type TwoInts struct {
	a, b int
}

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.String()[0:3]
}

func main() {
	// a := S{10}
	// a.print()
	// b := SType{20}
	// b.print()
	// c := SAlias{30}
	// c.print()
	// d := IntType(40)
	// d.print()
	// e := IntAlias(50)

	// two1 := new(TwoInts)
	// two1.a = 12
	// two1.b = 10
	// fmt.Println("sum:", two1.AddThem())
	// fmt.Println("param:", two1.AddToParan(20))
	// two2 := TwoInts{3, 4}
	// fmt.Println("sum:", two2.AddThem())

	m := myTime{time.Now()}
	fmt.Println(m.String())
	fmt.Println(m.first3Chars())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParan(param int) int {
	return tn.a + tn.b + param
}
