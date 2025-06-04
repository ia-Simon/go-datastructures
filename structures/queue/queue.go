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

type Queue[T any] struct {
	front *node[T]
	back  *node[T]
}

// Enqueue adds a value to the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	newNd := newNode(value)

	if q.front == nil {
		q.front = newNd
		q.back = newNd
		return
	}

	q.back.Next = newNd
	q.back = newNd
}

// Dequeue removes and returns the value at the front of
// the queue. A boolean is returned indicating whether
// value was filled from the queue (true), or the queue
// was empty (false).
func (q *Queue[T]) Dequeue() (value T, ok bool) {
	if q.front == nil {
		return value, false
	}

	dequeued := q.front
	q.front = q.front.Next
	if q.front == nil {
		q.back = nil
	}

	return dequeued.Value, true
}

// IsEmpty returns true if there are no values stored in the
// queue, false otherwise.
func (q *Queue[T]) IsEmpty() bool {
	return q.front == nil
}

// Front returns the value at the front of the queue, without
// removing it. A boolean is returned indicating whether
// value was filled from the queue (true), or the queue
// was empty (false).
func (q *Queue[T]) Front() (value T, ok bool) {
	if q.front == nil {
		return value, false
	}

	return q.front.Value, true
}

// Back returns the value at the back of the queue, without
// removing it. A boolean is returned indicating whether
// value was filled from the queue (true), or the queue
// was empty (false).
func (q *Queue[T]) Back() (value T, ok bool) {
	if q.front == nil {
		return value, false
	}

	return q.back.Value, true
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}
