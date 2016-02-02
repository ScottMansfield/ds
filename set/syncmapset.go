package set

import (
	"sync"
)

type SyncMapSet struct {
	lock  *sync.RWMutex
	items map[interface{}]struct{}
}

func NewSyncMapSet() *SyncMapSet {
	return &SyncMapSet{
		lock:  new(sync.RWMutex),
		items: make(map[interface{}]struct{}),
	}
}

func (s *SyncMapSet) Add(item interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items[item] = struct{}{}

	return nil
}

func (s *SyncMapSet) AddAll(items ...interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, item := range items {
		s.items[item] = struct{}{}
	}

	return nil
}

func (s *SyncMapSet) Contains(item interface{}) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.items[item]

	return ok, nil
}

func (s *SyncMapSet) ContainsAll(items ...interface{}) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, item := range items {
		if _, ok := s.items[item]; !ok {
			return false, nil
		}
	}

	return true, nil
}

func (s *SyncMapSet) Remove(item interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.items, item)

	return nil
}

func (s *SyncMapSet) RemoveAll(items ...interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, item := range items {
		delete(s.items, item)
	}

	return nil
}

func (s *SyncMapSet) Clear() error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = make(map[interface{}]struct{})

	return nil
}

func (s *SyncMapSet) Size() (int, error) {
	return len(s.items), nil
}

func (s *SyncMapSet) IsEmpty() (bool, error) {
	return (len(s.items) == 0), nil
}

func (s *SyncMapSet) ToSlice() ([]interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	ret := make([]interface{}, 0, len(s.items))
	for item, _ := range s.items {
		ret = append(ret, item)
	}

	return ret, nil
}
