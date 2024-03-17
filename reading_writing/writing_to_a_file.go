package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, err := os.OpenFile("output/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
