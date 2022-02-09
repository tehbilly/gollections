package gollections_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/tehbilly/gollections"
	"testing"
)

func TestOption(t *testing.T) {
	t.Run("Some", func(t *testing.T) {
		o := gollections.Some("foo")

		assert.True(t, o.IsSome())
		assert.False(t, o.IsNone())

		o.IfSome(func(v string) { assert.Equal(t, "foo", v) })
		o.IfNone(func() { assert.Fail(t, "should not be called") })

		assert.Equal(t, "foo", o.Get())
		assert.Equal(t, "foo", o.OrElse("other"))
	})

	t.Run("None", func(t *testing.T) {
		o := gollections.None[string]()

		assert.False(t, o.IsSome())
		assert.True(t, o.IsNone())

		o.IfSome(func(v string) { assert.Fail(t, "should not be called") })
		var called bool
		o.IfNone(func() { called = true })
		assert.True(t, called)

		assert.Panics(t, func() { o.Get() })
		assert.Equal(t, "other", o.OrElse("other"))
	})
}
