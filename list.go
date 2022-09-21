package gollections

type List[T comparable] struct {
	items []T
}

func (l *List[T]) Iterator() Iterator[T] {
	return &listIterator[T]{list: l}
}

func (l *List[T]) Size() int {
	return len(l.items)
}

func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

func (l *List[T]) Contains(o T) bool {
	for _, i := range l.items {
		if i == o {
			return true
		}
	}
	return false
}

func (l *List[T]) Add(o T) {
	l.items = append(l.items, o)
}

func (l *List[T]) Remove(o T) bool {
	pos := -1

	for i, t := range l.items {
		if t == o {
			pos = i
			break
		}
	}

	if pos == -1 {
		return false
	}

	l.items = append(l.items[:pos], l.items[pos+1:]...)

	return true
}

func (l *List[T]) AddAll(other Collection[T]) {
	oi := other.Iterator()

	for oi.HasNext() {
		l.items = append(l.items, oi.Next())
	}
}

type listIterator[T comparable] struct {
	list *List[T]
	pos  int
}

func (l *listIterator[T]) HasNext() bool {
	return l.pos < len(l.list.items)
}

func (l *listIterator[T]) Next() T {
	item := l.list.items[l.pos]
	l.pos += 1
	return item
}
