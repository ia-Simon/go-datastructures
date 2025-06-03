package stack

type node[T any] struct {
	Value T
	Next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		Value: value,
	}
}

type Stack[T any] struct {
	top *node[T]
}

// Push puts a value on the top of the stack.
func (s *Stack[T]) Push(value T) {
	newNd := newNode(value)

	if s.top == nil {
		s.top = newNd
		return
	}

	newNd.Next = s.top
	s.top = newNd
}

// Pop removes and returns the value on the top of the stack. A
// boolean is returned indicating whether value was filled from
// the stack (true), or the stack was empty (false).
func (s *Stack[T]) Pop() (value T, ok bool) {
	if s.top == nil {
		return value, false
	}

	popped := s.top
	s.top = s.top.Next

	return popped.Value, true
}

// Peek returns the value on the top of the stack, without removing
// it. A boolean is returned indicating whether value was filled from
// the stack (true), or the stack was empty (false).
func (s *Stack[T]) Peek() (value T, ok bool) {
	if s.top == nil {
		return value, false
	}

	return s.top.Value, true
}

// IsEmpty returns true if there are no values in the stack, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}
