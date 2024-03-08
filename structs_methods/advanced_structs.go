package main

import (
	"fmt"
	"unsafe"
)

type T1 struct {
	a, b int64
}

type T2 struct {
	t1 *T1
}

type T3 struct {
	t1 T1
}

type number struct {
	f float32
}

type nr number
type nrAlias = number

func main() {
	fmt.Println(unsafe.Sizeof(T1{}))
	fmt.Println(unsafe.Sizeof(T2{&T1{}}))
	fmt.Println(unsafe.Sizeof(T3{}))

	a := number{5.0}
	b := nr{5.0}
	c := nrAlias{5.0}

	var d number = number(b)
	var e = b
	fmt.Println(a, b, c, d, e)

}
