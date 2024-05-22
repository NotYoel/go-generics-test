package main

import (
	"fmt"

	"github.com/NotYoel/go-generics-test/ds"
)

func main() {
	lst := &ds.LinkedList[int]{}

	lst.Push(10)
	lst.Push(45)
	lst.Push(90)

	fmt.Println(lst.Remove(1), lst.All())
}
