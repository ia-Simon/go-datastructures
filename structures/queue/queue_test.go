package queue_test

import (
	"testing"

	"github.com/ia-Simon/go-datastrutures/structures/queue"
	"github.com/stretchr/testify/require"
)

func TestQueue_Enqueue(t *testing.T) {
	q := queue.New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var vals []int
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		vals = append(vals, val)
	}

	require.Equal(t, []int{1, 2, 3}, vals)
}

func TestQueue_Dequeue(t *testing.T) {
	q := queue.New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Dequeue()
	q.Dequeue()

	var vals []int
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		vals = append(vals, val)
	}

	require.Equal(t, []int{3}, vals)
}

func TestQueue_Front(t *testing.T) {
	q := queue.New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	val, ok := q.Front()
	require.True(t, ok)
	require.Equal(t, 1, val)

	var vals []int
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		vals = append(vals, val)
	}

	require.Equal(t, []int{1, 2}, vals)
}

func TestQueue_Back(t *testing.T) {
	q := queue.New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	val, ok := q.Back()
	require.True(t, ok)
	require.Equal(t, 2, val)

	var vals []int
	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		vals = append(vals, val)
	}

	require.Equal(t, []int{1, 2}, vals)
}

func TestQueue_IsEmpty(t *testing.T) {
	q := queue.New[int]()

	empty := q.IsEmpty()
	require.True(t, empty)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	empty = q.IsEmpty()
	require.False(t, empty)
}
