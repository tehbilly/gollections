package gollections_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tehbilly/gollections"
	"testing"
)

func TestConcurrentMap_Get(t *testing.T) {
	var c gollections.ConcurrentMap[string, string]
	c.Set("foo", "bar")

	assert.Equal(t, "bar", c.Get("foo"))

	v, ok := c.GetOK("foo")
	assert.Equal(t, "bar", v)
	assert.Equal(t, true, ok)

	v, ok = c.GetOK("baz")
	assert.Equal(t, "", v)
	assert.Equal(t, false, ok)
}

func TestConcurrentMap_Delete(t *testing.T) {
	var c gollections.ConcurrentMap[string, string]
	c.Set("foo", "bar")

	// Ensure value exists
	v, ok := c.GetOK("foo")
	assert.Equal(t, "bar", v)
	assert.Equal(t, true, ok)

	// Delete value that exists
	o, ok := c.Delete("foo")
	assert.Equal(t, "bar", o)
	assert.Equal(t, true, ok)

	// Ensure value was deleted
	v, ok = c.GetOK("foo")
	assert.Equal(t, "", v)
	assert.Equal(t, false, ok)

	// Delete value that doesn't exist
	o, ok = c.Delete("foo")
	assert.Equal(t, "", o)
	assert.Equal(t, false, ok)
}

func TestConcurrentMap_Keys(t *testing.T) {
	var c gollections.ConcurrentMap[string, string]

	c.Set("foo", "foo")
	c.Set("bar", "bar")
	c.Set("baz", "baz")

	assert.ElementsMatch(t, []string{"foo", "bar", "baz"}, c.Keys())
}

func TestConcurrentMap_ForEach(t *testing.T) {
	var c gollections.ConcurrentMap[string, string]

	c.Set("foo", "foo")
	c.Set("bar", "bar")
	c.Set("baz", "baz")

	t.Run("IterAll", func(t *testing.T) {
		var iters int
		c.ForEach(func(k string, v string) bool {
			iters += 1
			return true
		})

		// All keys were seen
		assert.Equal(t, len(c.Keys()), iters)
	})

	t.Run("EarlyBreak", func(t *testing.T) {
		var iters int
		c.ForEach(func(k string, v string) bool {
			iters += 1
			return false
		})

		// Only iterated once, even though three keys exist
		assert.Equal(t, 1, iters)
	})
}
