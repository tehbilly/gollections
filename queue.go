package gollections

import "sync"

type QueueOptions struct {
	Synchronized bool
	Capacity     int
}

type Queue[T any] struct {
	sync  bool
	mtx   sync.Mutex
	cap   int
	items []T
}

func (q *Queue[T]) Push(v T) bool {
	q.lock()
	defer q.unlock()

	if q.cap > 0 && len(q.items) >= q.cap {
		return false
	}
	q.items = append(q.items, v)
	return true
}

func (q *Queue[T]) Pop() Option[T] {
	q.lock()
	defer q.unlock()

	if len(q.items) == 0 {
		return None[T]()
	}

	var v T
	v, q.items = q.items[0], q.items[1:]
	return Some(v)
}

func (q *Queue[T]) Size() int {
	q.lock()
	size := len(q.items)
	q.unlock()
	return size
}

func (q *Queue[T]) lock() {
	if q.sync {
		q.mtx.Lock()
	}
}

func (q *Queue[T]) unlock() {
	if q.sync {
		q.mtx.Unlock()
	}
}

// NewQueue creates a new Queue instance, customized by QueueOptions
func NewQueue[T any](qo QueueOptions) *Queue[T] {
	return &Queue[T]{
		sync: qo.Synchronized,
		cap:  qo.Capacity,
	}
}
