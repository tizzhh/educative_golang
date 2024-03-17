package main

import (
	"flag"
	"fmt"
)

func main() {
	sptr := flag.String("lang", "Go", "a str")
	nptr := flag.Int("num", 108, "an int")
	bptr := flag.Bool("bool", false, "a bool")
	var str string
	flag.StringVar(&str, "str", "crystal", "a string var")
	flag.Parse()
	fmt.Println("lang:", *sptr)
	fmt.Println("num:", *nptr)
	fmt.Println("bool:", *bptr)
	fmt.Println("str:", str)
	fmt.Println("tail:", flag.Args())
}
