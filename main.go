package main

import (
	"fmt"

	"github.com/NotYoel/go-generics-test/ds"
)

func main() {
	// LINKED LIST
	lst := &ds.LinkedList[int]{}

	lst.Push(10)
	lst.Push(45)
	lst.Push(90)

	fmt.Println(lst.Remove(1), lst.All(), lst.Contains(90), lst.Contains(45)) // Output: <nil>, [10, 90], true, false

	// STACK
	s := &ds.Stack[int]{}

	s.Push(10)
	s.Push(20)

	p1, _ := s.Pop()
	p2, _ := s.Pop()
	p3, err := s.Pop()

	fmt.Println(s.All(), p1, p2, s.All(), p3, err) // Output: [20, 10], 20, 10, [], 0, cannot pop from empty stack
}
