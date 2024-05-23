package ds

import (
	"fmt"
	"reflect"
)

type Stack[T any] struct {
	head *Element[T]
}

func (s *Stack[T]) Head() *Element[T] {
	return s.head
}

func (s *Stack[T]) Push(itm T) {
	e := &Element[T]{
		next: s.head,
		val:  itm,
	}

	s.head = e
}

func (s *Stack[T]) Pop() (val T, err error) {
	if s.Length() == 0 {
		return *new(T), fmt.Errorf("cannot pop from empty stack")
	}

	val = s.head.val
	s.head = s.head.next

	return
}

func (lst *Stack[T]) Contains(itm T) bool {
	for curr := lst.head; curr != nil; curr = curr.next {
		if reflect.DeepEqual(curr.val, itm) {
			return true
		}
	}

	return false
}

func (s *Stack[T]) Length() (len int) {
	for curr := s.head; curr != nil; curr = curr.next {
		len++
	}

	return
}

func (s *Stack[T]) All() (itms []T) {
	for curr := s.head; curr != nil; curr = curr.next {
		itms = append(itms, curr.val)
	}

	return
}
