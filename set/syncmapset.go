package set

import (
    "sync"
)

type SyncMapSet struct {
    items map[interface{}]interface{}
    lock sync.RWMutex
}

func NewSyncMapSet() *SyncMapSet {
    items := make(map[interface{}]interface{})

    return &SyncMapSet{
        items: items,
    }
}

func (s *SyncMapSet) Add(item interface{}) error {
    s.items[item] = nil
    return nil
}
