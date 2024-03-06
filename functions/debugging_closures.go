package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	a := 2 + 5
	where()
	b := a + 100
	fmt.Println(a, b)
	where()
}
