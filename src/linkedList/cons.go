package linkedList

import "fmt"

type cons[T comparable] struct {
	value T
	next  node[T]
	prev  node[T]
}

func (c *cons[T]) prepend(val T) node[T] {
	newHead := &cons[T]{
		value: val,
		next:  c,
		prev:  newNode[T](),
	}

	c.prev = newHead

	return newHead
}

func (c *cons[T]) append(val T) node[T] {
	newTail := &cons[T]{
		value: val,
		next:  newNode[T](),
		prev:  c,
	}

	c.next = newTail

	return newTail
}

func (c cons[T]) val() T {
	return c.value
}

func (c cons[T]) contains(val T) bool {
	return c.value == val || c.next.contains(val)
}

func (c *cons[T]) removeNext() {
	c.next = newNode[T]()
}

func (c *cons[T]) removePrev() {
	c.prev = newNode[T]()
}

func (c cons[T]) String() string {
	return fmt.Sprintf("%+v -> %s", c.value, c.next)
}

func (c cons[T]) stringInverse() string {
	return fmt.Sprintf("%+v <- %s", c.value, c.prev.stringInverse())
}
