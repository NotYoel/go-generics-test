package ds

type Element[T any] struct {
	val  T
	next *Element[T]
}

func (e *Element[T]) Value() T {
	return e.val
}

func (e *Element[T]) Next() *Element[T] {
	return e.next
}
