package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dstName, srcName string) (writter int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	copied, _ := CopyFile("output/target.txt", "products.txt")
	fmt.Println("copy done", copied)
}
