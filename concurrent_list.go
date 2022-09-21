package gollections

import "sync"

type ConcurrentList[T comparable] struct {
	mtx  sync.RWMutex
	list List[T]
}

func (c *ConcurrentList[T]) Iterator() Iterator[T] {
	return &concurrentListIterator[T]{list: c}
}

func (c *ConcurrentList[T]) Size() int {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.list.Size()
}

func (c *ConcurrentList[T]) IsEmpty() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.list.IsEmpty()
}

func (c *ConcurrentList[T]) Contains(o T) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.list.Contains(o)
}

func (c *ConcurrentList[T]) Add(o T) {
	c.mtx.Lock()
	c.list.Add(o)
	c.mtx.Unlock()
}

func (c *ConcurrentList[T]) Remove(o T) bool {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.list.Remove(o)
}

func (c *ConcurrentList[T]) AddAll(other Collection[T]) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.list.AddAll(other)
}

type concurrentListIterator[T comparable] struct {
	list *ConcurrentList[T]
	pos  int
}

func (l *concurrentListIterator[T]) HasNext() bool {
	l.list.mtx.RLock()
	defer l.list.mtx.RUnlock()
	return l.pos < len(l.list.list.items)
}

func (l *concurrentListIterator[T]) Next() T {
	l.list.mtx.Lock()
	defer l.list.mtx.Unlock()

	item := l.list.list.items[l.pos]
	l.pos += 1
	return item
}
