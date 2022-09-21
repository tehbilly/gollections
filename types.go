package gollections

type Iterator[T comparable] interface {
	HasNext() bool
	Next() T
}

type Iterable[T comparable] interface {
	Iterator() Iterator[T]
}

type Collection[T comparable] interface {
	Iterable[T]

	Size() int
	IsEmpty() bool

	Contains(o T) bool
	Add(o T)
	Remove(o T) bool
	AddAll(other Collection[T])
}
