package ds

import (
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

	var val T

	for curr, i := lst.head, 0; curr != nil; curr = curr.next {
		if i == index {
			val = curr.val
			break
		}
		i++
	}

	return val, nil
}

func (lst *LinkedList[T]) Remove(index int) error {
	// Index out of bounds
	if lst.indexOutOfBounds(index) {
		return fmt.Errorf("index %d not found", index)
	}

	if index == 0 {
		if lst.Length() == 1 {
			lst.head, lst.tail = nil, nil
		} else {
			lst.head = lst.head.next
		}

		return nil
	} else if index == lst.Length()-1 {
		if lst.Length() == 2 {
			lst.head.next, lst.tail = nil, lst.head
			return nil
		}
	}

	for curr, i, last := lst.head, 0, lst.head; curr != nil; curr = curr.next {
		if i == index {
			last.next = curr.next

			if index == lst.Length() {
				lst.tail = last
			}
			break
		}

		i++
		last = curr

	}

	return nil
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
