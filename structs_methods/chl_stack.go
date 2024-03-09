package main

import (
	"fmt"
)

const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix < LIMIT {
		st.data[st.ix] = n
		st.ix++
	}
}

func (st *Stack) Pop() int {
	if st.ix == 0 {
		return -1
	}
	st.ix--
	val := st.data[st.ix]
	st.data[st.ix] = 0
	return val
}

func (st *Stack) String() string {
	var res string
	for i := 0; i < st.ix; i++ {
		res += fmt.Sprintf("[%d:%d]", i, st.data[i])
	}
	return res
}

func main() {
	st := new(Stack)
	st.Push(1)
	st.Push(5)
	st.Push(4)
	st.Push(4)
	fmt.Println(st)
	st.Pop()
	st.Pop()
	st.Pop()
	fmt.Println(st)
}
