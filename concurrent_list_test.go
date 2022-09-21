package gollections_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehbilly/gollections"
)

func TestConcurrentList(t *testing.T) {
	var tl gollections.ConcurrentList[string]

	assert.Equal(t, 0, tl.Size())
	assert.False(t, tl.Contains("foo"))
	assert.True(t, tl.IsEmpty())

	tl.Add("foo")
	tl.Add("bar")
	tl.Add("baz")

	assert.Equal(t, 3, tl.Size())
	assert.True(t, tl.Contains("foo"))
	assert.False(t, tl.IsEmpty())

	tli := tl.Iterator()
	seen := 0
	assert.True(t, tli.HasNext())

	for tli.HasNext() {
		seen += 1
		switch seen {
		case 1:
			assert.Equal(t, "foo", tli.Next())
		case 2:
			assert.Equal(t, "bar", tli.Next())
		case 3:
			assert.Equal(t, "baz", tli.Next())
		default:
			assert.Fail(t, "Unexpected item!")
		}
	}

	assert.Equal(t, 3, seen)
	assert.False(t, tli.HasNext())
}

func TestConcurrentList_AddAll(t *testing.T) {
	var l1 gollections.ConcurrentList[string]
	l1.Add("foo")
	l1.Add("bar")
	l1.Add("baz")
	assert.Equal(t, 3, l1.Size())

	var l2 gollections.ConcurrentList[string]
	l2.Add("one")
	l2.Add("two")
	l2.Add("three")
	assert.Equal(t, 3, l2.Size())

	l1.AddAll(&l2)

	assert.Equal(t, 6, l1.Size())
}

func TestConcurrentList_Remove(t *testing.T) {
	var tl gollections.ConcurrentList[string]
	tl.Add("foo")
	tl.Add("bar")
	tl.Add("baz")
	assert.Equal(t, 3, tl.Size())

	assert.False(t, tl.Remove("not present"))
	assert.Equal(t, 3, tl.Size())

	assert.True(t, tl.Remove("foo"))
	assert.Equal(t, 2, tl.Size())
}
