package queue

type node[T any] struct {
	Value T
	Next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		Value: value,
	}
}
