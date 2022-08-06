package linkedList

import "fmt"

type node[T comparable] interface {
	fmt.Stringer
	val() T
	prepend(T) node[T]
	append(T) node[T]
	contains(T) bool
	removePrev()
	removeNext()
	stringInverse() string
}

func newNode[T comparable]() node[T] {
	return none[T]{}
}

type LinkedList[T comparable] interface {
	fmt.Stringer
	Prepend(T)
	Append(T)
	Pop() (T, error)
	Contains(T) bool
	StringInverse() string
}

type List[T comparable] struct {
	head  node[T]
	tail  node[T]
	count int
}

func New[T comparable]() LinkedList[T] {
	return &List[T]{
		head:  newNode[T](),
		tail:  newNode[T](),
		count: 0,
	}
}

func (l *List[T]) Prepend(val T) {
	l.head = l.head.prepend(val)

	if l.count == 0 {
		l.tail = l.head
	}

	l.count += 1
}

func (l *List[T]) Append(val T) {
	l.tail = l.tail.append(val)

	if l.count == 0 {
		l.head = l.tail
	}

	l.count += 1
}

func (l *List[T]) Pop() (T, error) {
	if l.count < 1 {
		return l.head.val(), fmt.Errorf("the list is empty")
	}

	head := l.head
	switch c := head.(type) {
	case *cons[T]:
		c.next.removePrev()
		l.head = c.next
	}

	return head.val(), nil
}

func (l *List[T]) Contains(val T) bool {
	return l.head.contains(val)
}

func (l List[T]) String() string {
	s := l.head.String()

	return s[:len(s)-4]
}

func (l List[T]) StringInverse() string {
	s := l.tail.stringInverse()

	return s[:len(s)-4]
}
