package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("input.data")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		if err == io.EOF {
			return
		}
		fmt.Println("file string:", inputString)
	}
}
