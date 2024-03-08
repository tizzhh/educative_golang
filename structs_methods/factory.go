package main

import "fmt"

type File struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

type Foo map[string]string

type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	f := NewFile(10, "../test.txt")
	fmt.Println(f)

	y := new(Bar)
	y.thingOne = "hello"
	y.thingTwo = 1

	// z := make(Bar) // compile err
	// z.thingOne = "hello"
	// z.thingTwo = 1

	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// panic
	// u := new(Foo)
	// (*u)["x"] = "goodbye"
}
