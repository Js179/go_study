package structure

import "sync"

type LockMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// Put map的put方法
func (m *LockMap[K, V]) Put(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data == nil {
		m.data = make(map[K]V)
	}
	m.data[k] = v
}

// Get map的get方法，通过key获取指定的value
func (m *LockMap[K, V]) Get(k K) (v V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok = m.data[k]
	return
}

// Delete map的delete方法，通过key删除指定项
func (m *LockMap[K, V]) Delete(k K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, k)
}

// Len 返回map的长度
func (m *LockMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}
