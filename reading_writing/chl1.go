package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrword, nrlines int

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('S')
	if err != nil {
		fmt.Println(err)
		return
	}
	Counters(input)
	fmt.Println(nrchars, nrword, nrlines)
}

func Counters(input string) {
	nrchars = len(input[:len(input)-1]) - strings.Count(input, "\n")
	nrword = len(strings.Split(input, " "))
	nrlines = strings.Count(input, "\n") + 1
}
