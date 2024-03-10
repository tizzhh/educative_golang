package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.4
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
	fmt.Println(reflect.ValueOf(x).Kind() == reflect.Float64)

	var y float64 = 3.5
	fmt.Println(reflect.TypeOf(y))
	v := reflect.ValueOf(y)
	fmt.Println(v, v.Type(), v.Kind(), v.Float(), v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	z := v.Interface().(float64)
	fmt.Println(z)
	fmt.Println(v.CanSet())
	fmt.Println()
	v = reflect.ValueOf(&x)
	fmt.Println(v.Type(), v.CanSet())
	v = v.Elem()
	fmt.Println(v, v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface(), v)
}
