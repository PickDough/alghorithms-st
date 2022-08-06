package linkedList

type none[T comparable] struct{}

func (n none[T]) prepend(val T) node[T] {
	return &cons[T]{
		value: val,
		next:  newNode[T](),
		prev:  newNode[T](),
	}
}

func (n none[T]) append(val T) node[T] {
	return &cons[T]{
		value: val,
		next:  newNode[T](),
		prev:  newNode[T](),
	}
}

func (none[T]) val() T {
	var v T
	return v
}

func (none[T]) contains(val T) bool {
	return false
}

func (none[T]) removeNext() {

}

func (none[T]) removePrev() {

}

func (n none[T]) String() string {
	return ""
}

func (n none[T]) stringInverse() string {
	return ""
}
