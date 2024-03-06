package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	where := log.Print

	where()
	a := 2 * 5
	where()

	b := a + 100
	fmt.Println(a, b)
	where()

}
