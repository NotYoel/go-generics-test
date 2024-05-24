package ds

import "fmt"

type Queue[T any] struct {
	head, tail *Element[T]
}

func (q *Queue[T]) Enqueue(itm T) {
	e := &Element[T]{
		val: itm,
	}

	if q.tail == nil {
		q.head = e
		q.tail = e
	} else {
		q.tail.next = e
		q.tail = e
	}
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Length() == 0 {
		return *new(T), fmt.Errorf("cannot dequeue from empty queue")
	}

	val := q.head.val

	if q.Length() == 1 {
		q.head, q.tail = nil, nil
	} else {
		q.head = q.head.next
	}

	return val, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Length() == 0 {
		return *new(T), fmt.Errorf("cannot peek empty queue")
	}

	return q.head.val, nil
}

func (q *Queue[T]) Length() (len int) {
	for curr := q.head; curr != nil; curr = curr.next {
		len++
	}

	return
}

func (q *Queue[T]) All() (itms []T) {
	for curr := q.head; curr != nil; curr = curr.next {
		itms = append(itms, curr.val)
	}

	return
}
