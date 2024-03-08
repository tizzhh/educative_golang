package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    `How much there are`
}

type T struct {
	a int "This is a tag"
	b int `a raw stirng tag`
	c int `key1:"value1" key2:"value2"`
}

func main() {
	// tt := TagType{true, "barack", 1}
	// for i := 0; i < 2; i++ {
	// 	refTag(tt, i)
	// }

	t := T{}
	fmt.Println(reflect.TypeOf(t).Field(0).Tag)
	if field, ok := reflect.TypeOf(t).FieldByName("b"); ok {
		fmt.Println(field.Tag)
	}
	if field, ok := reflect.TypeOf(t).FieldByName("c"); ok {
		fmt.Println(field.Tag)
		fmt.Println(field.Tag.Get("key1"))
	}
	if field, ok := reflect.TypeOf(t).FieldByName("d"); ok {
		fmt.Println(field.Tag)
	} else {
		fmt.Println("Field not found")
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
