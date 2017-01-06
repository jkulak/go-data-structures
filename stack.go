package stack

import "fmt"

// Stack is an implementation of Stack data structure
type Stack []int

// New makes a new Stack
func New() Stack {
	return make(Stack, 0)
}

// Push method puts the element on top of the Stack
func (s Stack) Push(v int) Stack {
	return append(s, v)
}

// Pop method removes and returns the last element from the Stack
func (s Stack) Pop() (Stack, int) {
	l := len(s)
	if l == 0 {
		panic(fmt.Sprintf("Stack %v is empty, you can't pop from it", s))
	}
	return s[:l-1], s[l-1]
}

// Helper functions

// Populate populates the Stack with give array of integers
func (s Stack) Populate(values []int) Stack {
	return values
}

// Print prints the Stack to the stdout
func (s Stack) Print() {
	fmt.Printf("Printing the stack %v\n", s)
	for e := range s {
		fmt.Println(s[e])
	}
	fmt.Println()
}
