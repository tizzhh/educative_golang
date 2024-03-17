package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 | 5212 | Go"
	format                 = "%f | %d | %s"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {
	// fmt.Println("name:")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Println(firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println(f, i, s)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("input something")
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("input was:", input2[:len(input2)-1])
	}
}
