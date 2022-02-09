package gollections

import "sync"

// ConcurrentMap allows safe concurrent access to an underlying map[K]V
type ConcurrentMap[K comparable, V any] struct {
	mtx   sync.RWMutex
	items map[K]V
}

func (m *ConcurrentMap[K, V]) ensureMap() {
	if m.items == nil {
		m.items = make(map[K]V)
	}
}

// Set sets the V value for the specified K key
func (m *ConcurrentMap[K, V]) Set(k K, v V) {
	m.mtx.Lock()
	m.ensureMap()
	m.items[k] = v
	m.mtx.Unlock()
}

// Get returns the V value for the specified K key, or the zero value if the key is not present
func (m *ConcurrentMap[K, V]) Get(k K) V {
	m.mtx.RLock()
	m.ensureMap()
	val := m.items[k]
	m.mtx.RUnlock()
	return val
}

// GetOK returns the V value and a boolean indicating whether the K key is present or not
func (m *ConcurrentMap[K, V]) GetOK(k K) (V, bool) {
	m.mtx.RLock()
	m.ensureMap()
	val, ok := m.items[k]
	m.mtx.RUnlock()
	return val, ok
}

// Delete removes an item from the map with the specified key. The V value (if any) and a boolean indicating whether the
// item existed or not is returned.
func (m *ConcurrentMap[K, V]) Delete(k K) (V, bool) {
	m.mtx.Lock()
	m.ensureMap()
	val, ok := m.items[k]
	if ok {
		delete(m.items, k)
	}
	m.mtx.Unlock()
	return val, ok
}

// Keys will return []K with all keys present in the map at the time of the call.
func (m *ConcurrentMap[K, V]) Keys() []K {
	var r []K

	m.mtx.RLock()
	m.ensureMap()
	for k := range m.items {
		r = append(r, k)
	}
	m.mtx.RUnlock()

	return r
}

// ForEach will iterate over key/value pairs in the map. This is not guaranteed to be stable as it does not operate on
// a snapshot of the underlying data. All keys present at the time ForEach is called will be iterated over, but the
// provided function will only be invoked if the value is present at the time of invocation. Any items added after
// iteration starts will not be seen.
func (m *ConcurrentMap[K, V]) ForEach(f func(k K, v V) bool) {
	keys := m.Keys()

	for _, k := range keys {
		v, ok := m.GetOK(k)
		if ok && !f(k, v) {
			break
		}
	}
}
