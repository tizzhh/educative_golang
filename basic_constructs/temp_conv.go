package main

import "fmt"

type Celcius float32
type Fahrenheit float32

func toFahrenheit(t Celcius) Fahrenheit {
	return Fahrenheit(t*9/5 + 32)
}

func main() {
	fmt.Println(toFahrenheit(100))
}
