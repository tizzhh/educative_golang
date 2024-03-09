package main

import (
	"fmt"
)

type employee struct {
	salary float32
}

func (e *employee) giveRaise(pct float32) {
	e.salary += e.salary * pct
}
