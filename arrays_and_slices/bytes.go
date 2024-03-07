package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.WriteString("ABC")
	b.WriteString("DEF")
	fmt.Fprintf(&b, " A  num %d a string %v\n", 10, "aboba")
	b.WriteString("DONE")
	fmt.Println(b.String())
}
