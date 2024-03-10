package mystack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) isEmpty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack Stack) Top() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("empty stack")
	}
	return stack[stack.Len()-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.isEmpty() {
		return nil, errors.New("empty stack")
	}
	val := (*stack)[stack.Len()-1]
	*stack = (*stack)[:stack.Len()-1]
	return val, nil
}
