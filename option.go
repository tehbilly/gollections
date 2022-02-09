package gollections

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	IfSome(f func(v T))
	IfNone(f func())
	OrElse(e T) T
	Get() T
}

type option[T any] struct {
	isSome bool
	v      T
}

func (o option[T]) IsSome() bool {
	return o.isSome
}

func (o option[T]) IsNone() bool {
	return !o.isSome
}

func (o option[T]) IfSome(f func(v T)) {
	if o.isSome {
		f(o.v)
	}
}

func (o option[T]) IfNone(f func()) {
	if !o.isSome {
		f()
	}
}

func (o option[T]) OrElse(e T) T {
	if o.isSome {
		return o.v
	}
	return e
}

func (o option[T]) Get() T {
	if o.isSome {
		return o.v
	}
	panic("cannot get from None option")
}

func None[T any]() Option[T] {
	return option[T]{
		isSome: false,
	}
}

func Some[T any](v T) Option[T] {
	return option[T]{
		isSome: true,
		v:      v,
	}
}
