package main

import (
	"fmt"
	"math"
)

type polar struct {
	radius, angle float64
}

type cartesian struct {
	x, y float64
}

func main() {
	questions := make(chan polar)
	answers := createSolver(questions)
	go interact(questions, answers)
	fmt.Println("Answer:", <-answers)
}

func interact(questions chan polar, answers chan cartesian) {
	fmt.Println("Input radius, angle:")
	var radius, angle float64
	fmt.Scan(&radius, &angle)
	questions <- polar{radius, angle}
}

func createSolver(questions chan polar) chan cartesian {
	cart := cartesian{}
	answers := make(chan cartesian)
	go func() {
		pol := <-questions
		cart.x = pol.radius * math.Cos(float64(pol.angle)*math.Pi/180.0)
		cart.y = pol.radius * math.Sin(float64(pol.angle)*math.Pi/180.0)
		answers <- cart
	}()
	return answers
}
