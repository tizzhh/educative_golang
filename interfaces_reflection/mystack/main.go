package main

import (
	"fmt"
	"mystack/mystack"
)

func main() {
	st := mystack.Stack{}
	fmt.Println(st)
	st.Push(1)
	st.Push(2)
	st.Push(3)
	fmt.Println(st)
	val, _ := st.Pop()
	st.Push("aboba")
	st.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	fmt.Println(val, st)
}
