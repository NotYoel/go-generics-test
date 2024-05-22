package main

import (
	"fmt"

	"github.com/NotYoel/go-generics-test/ds"
)

func main() {
	lst := &ds.LinkedList[int]{}
	lst.Push(10)
	lst.Push(30)

	fmt.Println(lst.All())
}
