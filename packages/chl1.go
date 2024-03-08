package main

import (
	"container/list"
	"fmt"
)

func insertListElements(n int) *list.List {
	lst := list.New()
	for i := 1; i <= n; i++ {
		lst.PushBack(i)
	}
	return lst
}

func main() {
	n := 5
	resLst := insertListElements(n)
	for ptr := resLst.Front(); ptr != nil; ptr = ptr.Next() {
		fmt.Println(ptr.Value)
	}
}
