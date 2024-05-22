package ds

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head, tail *Element[T]
}

func (lst *LinkedList[T]) Head() *Element[T] {
	return lst.head
}

func (lst *LinkedList[T]) Tail() *Element[T] {
	return lst.tail
}

func (lst *LinkedList[T]) Push(v T) {
	e := &Element[T]{
		val: v,
	}

	// List is empty
	if lst.tail == nil {
		lst.head = e
		lst.tail = e
	} else {
		lst.tail.next = e
		lst.tail = e
	}
}

func (lst *LinkedList[T]) Get(index int) (T, error) {
	// Index out of bounds
	if lst.indexOutOfBounds(index) {
		// *new(T) is how you can return the zero value for generic types.
		return *new(T), fmt.Errorf("index %d not found", index)
	}

	for curr, i := lst.head, 0; curr != nil; curr = curr.next {
		if i == index {
			return curr.val, nil
		}
		i++
	}

	return *new(T), errors.New("unknown error")
}

func (lst *LinkedList[T]) Length() (len int) {
	for curr := lst.head; curr != nil; curr = curr.next {
		len++
	}

	return
}

func (lst *LinkedList[T]) All() (itms []T) {
	for curr := lst.head; curr != nil; curr = curr.next {
		itms = append(itms, curr.val)
	}

	return
}

func (lst *LinkedList[T]) indexOutOfBounds(index int) bool {
	return index >= lst.Length() || index < 0
}
