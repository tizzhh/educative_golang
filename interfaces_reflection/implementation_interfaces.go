package main

import "fmt"

type List []int

func (l List) Len() int { return len(l) }

func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	var lst List
	// CountInto(lst, 1, 10) // since l *List Append
	if LongEnough(lst) {
		fmt.Println("Long enough")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Println("this time long enough", plst)
	}
}
