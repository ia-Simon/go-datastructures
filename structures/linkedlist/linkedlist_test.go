package linkedlist_test

import (
	"testing"

	"github.com/ia-Simon/go-datastrutures/structures/linkedlist"

	"github.com/stretchr/testify/require"
)

func TestLinkedList_InsertFront(t *testing.T) {
	ll := linkedlist.New[int]()

	ll.InsertFront(1)
	ll.InsertFront(4)
	ll.InsertFront(3)

	var vals []int
	for i := range ll.Iter() {
		vals = append(vals, i)
	}

	require.Equal(t, []int{3, 4, 1}, vals)
}

func TestLinkedList_InsertEnd(t *testing.T) {
	ll := linkedlist.New[int]()

	ll.InsertEnd(1)
	ll.InsertEnd(4)
	ll.InsertEnd(3)

	var vals []int
	for i := range ll.Iter() {
		vals = append(vals, i)
	}

	require.Equal(t, []int{1, 4, 3}, vals)
}

func TestLinkedList_InsertAt(t *testing.T) {
	t.Run("nominal use", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(3)
		ll.InsertAt(2, 1)

		var vals []int
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{1, 2, 3}, vals)
	})

	t.Run("insert above list end", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(2)
		ll.InsertEnd(3)
		ll.InsertAt(10, 100)

		var vals []int
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{1, 2, 3, 10}, vals)
	})
}

func TestLinkedList_DeleteFront(t *testing.T) {
	t.Run("nominal use", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(2)
		ll.InsertEnd(3)
		ll.DeleteFront()

		var vals []int
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{2, 3}, vals)
	})

	t.Run("delete when head only", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.DeleteFront()

		var vals = make([]int, 0)
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{}, vals)
	})

	t.Run("delete when empty", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.DeleteFront()

		var vals = make([]int, 0)
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{}, vals)
	})
}

func TestLinkedList_DeleteEnd(t *testing.T) {
	t.Run("nominal use", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(2)
		ll.InsertEnd(3)
		ll.DeleteEnd()

		var vals []int
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{1, 2}, vals)
	})

	t.Run("delete when head only", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.DeleteEnd()

		var vals = make([]int, 0)
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{}, vals)
	})

	t.Run("delete when empty", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.DeleteEnd()

		var vals = make([]int, 0)
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{}, vals)
	})
}

func TestLinkedList_DeleteAt(t *testing.T) {
	t.Run("nominal use", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(2)
		ll.InsertEnd(3)
		ll.DeleteAt(1)

		var vals []int
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{1, 3}, vals)
	})

	t.Run("delete above list size", func(t *testing.T) {
		ll := linkedlist.New[int]()

		ll.InsertFront(1)
		ll.InsertEnd(2)
		ll.InsertEnd(3)
		ll.DeleteAt(5)

		var vals = make([]int, 0)
		for i := range ll.Iter() {
			vals = append(vals, i)
		}

		require.Equal(t, []int{1, 2, 3}, vals)
	})
}
