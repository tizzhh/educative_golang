package main

import "fmt"

func main() {
	var i1 int
	var f1 float32
	i1, _, f1 = threeVals()
	fmt.Println(i1, f1)

	n := 0
	reply := &n
	fmt.Println("before ", *reply)
	multiply(5, 10, reply)
	fmt.Println("after ", *reply)
	fmt.Println("after ", n)
	greeting("aboba", "privet", "hey")
}

func threeVals() (int, int, float32) {
	return 5, 6, 7.5
}

func multiply(a, b int, reply *int) {
	*reply = a * b
}

func greeting(who ...string) {
	for _, v := range who {
		fmt.Println("word is ", v)
	}
}
