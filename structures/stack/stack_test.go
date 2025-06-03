package stack_test

import (
	"testing"

	"github.com/ia-Simon/go-datastrutures/structures/stack"

	"github.com/stretchr/testify/require"
)

func TestStack_Push(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	var vals []int
	for !s.IsEmpty() {
		val, _ := s.Pop()
		vals = append(vals, val)
	}

	require.Equal(t, []int{3, 2, 1}, vals)
}

func TestStack_Pop(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, ok := s.Pop()
	require.True(t, ok)
	require.Equal(t, 3, val)

	var vals []int
	for !s.IsEmpty() {
		val, _ := s.Pop()
		vals = append(vals, val)
	}

	require.Equal(t, []int{2, 1}, vals)
}

func TestStack_Peek(t *testing.T) {
	s := stack.New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, ok := s.Peek()
	require.True(t, ok)
	require.Equal(t, 3, val)

	var vals []int
	for !s.IsEmpty() {
		val, _ := s.Pop()
		vals = append(vals, val)
	}

	require.Equal(t, []int{3, 2, 1}, vals)
}

func TestStack_IsEmpty(t *testing.T) {
	s := stack.New[int]()

	empty := s.IsEmpty()
	require.True(t, empty)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	empty = s.IsEmpty()
	require.False(t, empty)
}
