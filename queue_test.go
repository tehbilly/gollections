package gollections_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tehbilly/gollections"
	"testing"
)

func TestQueue(t *testing.T) {
	var q gollections.Queue[string]

	t.Run("Push", func(t *testing.T) {
		assert.True(t, q.Push("one"))
		assert.True(t, q.Push("two"))
		assert.True(t, q.Push("three"))
	})

	t.Run("Pop", func(t *testing.T) {
		assert.Equal(t, "one", q.Pop().Get())
		assert.Equal(t, "two", q.Pop().Get())
		assert.Equal(t, "three", q.Pop().Get())
		assert.True(t, q.Pop().IsNone())
	})

	t.Run("Size", func(t *testing.T) {
		assert.Equal(t, 0, q.Size())
		assert.True(t, q.Push("one"))
		assert.Equal(t, 1, q.Size())
		assert.Equal(t, "one", q.Pop().Get())
		assert.Equal(t, 0, q.Size())
	})
}

func TestQueue_Bounded(t *testing.T) {
	q := gollections.NewQueue[string](gollections.QueueOptions{Capacity: 2})

	t.Run("Push", func(t *testing.T) {
		assert.True(t, q.Push("one"))
		assert.True(t, q.Push("two"))
		assert.False(t, q.Push("three"))
	})

	t.Run("Pop", func(t *testing.T) {
		assert.Equal(t, "one", q.Pop().Get())
		assert.Equal(t, "two", q.Pop().Get())
		assert.True(t, q.Pop().IsNone())
	})

	t.Run("Size", func(t *testing.T) {
		assert.Equal(t, 0, q.Size())
		assert.True(t, q.Push("one"))
		assert.Equal(t, 1, q.Size())
		assert.True(t, q.Push("two"))
		assert.Equal(t, 2, q.Size())
		assert.False(t, q.Push("three"))
		assert.Equal(t, 2, q.Size())
	})
}

func TestQueue_Synchronized(t *testing.T) {
	q := gollections.NewQueue[string](gollections.QueueOptions{Synchronized: true})

	t.Run("Push", func(t *testing.T) {
		assert.True(t, q.Push("one"))
		assert.True(t, q.Push("two"))
		assert.True(t, q.Push("three"))
	})

	t.Run("Pop", func(t *testing.T) {
		assert.Equal(t, "one", q.Pop().Get())
		assert.Equal(t, "two", q.Pop().Get())
		assert.Equal(t, "three", q.Pop().Get())
	})

	t.Run("Size", func(t *testing.T) {
		assert.Equal(t, 0, q.Size())
		assert.True(t, q.Push("one"))
		assert.Equal(t, 1, q.Size())
		assert.True(t, q.Push("two"))
		assert.Equal(t, 2, q.Size())
		assert.True(t, q.Push("three"))
		assert.Equal(t, 3, q.Size())
	})
}
