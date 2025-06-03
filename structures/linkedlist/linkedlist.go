package linkedlist

import (
	"iter"
)

type node[T any] struct {
	Value T
	Next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		Value: value,
	}
}

type LinkedList[T any] struct {
	head *node[T]
}

// InsertFront inserts a value at the beginning (head) of the linked list.
func (ll *LinkedList[T]) InsertFront(value T) {
	newNd := newNode(value)

	if ll.head == nil {
		ll.head = newNd
		return
	}

	newNd.Next = ll.head
	ll.head = newNd
}

// InsertEnd inserts a value at the end of the linked list.
func (ll *LinkedList[T]) InsertEnd(value T) {
	newNd := newNode(value)

	if ll.head == nil {
		ll.head = newNd
		return
	}

	current := ll.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNd
}

// InsertAt insert a value at the indicated position of the linked list or
// at the end, whichever comes first.
//
// To insert in the front of the list, refer to LinkedList.InsertFront. Any position
// lesser than or equal to 0 will be considered as 1.
func (ll *LinkedList[T]) InsertAt(value T, position int) {
	newNd := newNode(value)

	if ll.head == nil {
		ll.head = newNd
		return
	}

	if position <= 0 {
		position = 1
	}

	current := ll.head
	for i := 1; i < position && current.Next != nil; i++ {
		current = current.Next
	}

	newNd.Next = current.Next
	current.Next = newNd
}

// DeleteFront deletes the first value (head) of the linked list.
func (ll *LinkedList[T]) DeleteFront() {
	if ll.head == nil {
		return
	}

	ll.head = ll.head.Next
}

// DeleteEnd deletes the last value of the linked list.
func (ll *LinkedList[T]) DeleteEnd() {
	if ll.head == nil {
		return
	}

	if ll.head.Next == nil {
		ll.head = nil
		return
	}

	var last *node[T]
	current := ll.head
	for current.Next != nil {
		last = current
		current = current.Next
	}

	last.Next = nil
}

// DeleteAt deletes a value at the indicated position, if it exists.
//
// To delete from the front of the list, refer to LinkedList.DeleteFront. Any position
// lesser than or equal to 0 will be considered as 1.
func (ll *LinkedList[T]) DeleteAt(position int) {
	if ll.head == nil {
		return
	}

	if position <= 0 {
		position = 1
	}

	var last *node[T]
	current := ll.head
	for i := 0; i < position; i++ {
		last = current
		current = current.Next
		if current == nil {
			return
		}
	}

	last.Next = current.Next
}

// Iter allows range-based iteration over all the values in the linked list.
func (ll *LinkedList[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := ll.head
		for current != nil {
			if !yield(current.Value) {
				return
			}
			current = current.Next
		}
	}

}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
