package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	orig := "ABC"

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an int - exiting with %v\n", an, err)
		os.Exit(1)
	}
	fmt.Printf("int %d\n", an)
}
