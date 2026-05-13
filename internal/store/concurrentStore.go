package store

import (
	"sort"
	"sync"
)

// concurrentStore is a thread safe implementation of the Store interface.
type concurrentStore struct {
	mtx  sync.RWMutex
	data map[string]string
}

// NewConcurrentStore creates a new instance of the concurrent store with no data.
func NewConcurrentStore() *concurrentStore {
	return &concurrentStore{
		mtx:  sync.RWMutex{},
		data: make(map[string]string),
	}
}

func (s *concurrentStore) Set(key string, val string) {
	s.mtx.Lock()
	s.data[key] = val
	s.mtx.Unlock()
}

func (s *concurrentStore) Get(key string) (string, bool) {
	s.mtx.RLock()

	result, exists := s.data[key]

	s.mtx.RUnlock()

	return result, exists
}

func (s *concurrentStore) Delete(key string) bool {
	s.mtx.Lock()
	_, exists := s.data[key]
	delete(s.data, key)
	s.mtx.Unlock()
	return exists
}

func (s *concurrentStore) Exists(key string) bool {
	s.mtx.RLock()

	_, exists := s.data[key]

	s.mtx.RUnlock()

	return exists
}

func (s *concurrentStore) Keys() []string {
	s.mtx.RLock()

	var keys []string
	for k := range s.data {
		keys = append(keys, k)
	}

	s.mtx.RUnlock()

	sort.Strings(keys)

	return keys
}
