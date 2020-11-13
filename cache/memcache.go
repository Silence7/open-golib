package cache

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"sync"
)

type Proc func(key interface{}, value interface{})
type UpdateCb func(exist bool, valueInMap interface{}, newValue interface{}) interface{}

type MemCache interface {
	Set(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Upsert(key interface{}, value interface{}, cb UpdateCb)
	Delete(key interface{})
	Foreach(f Proc)
}

type SyncMap struct {
	Cache sync.Map
	mu sync.Mutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		Cache: sync.Map{},
	}
}

func (m *SyncMap) Set(key interface{}, value interface{}) {
	m.Cache.Store(key, value)
}

func (m *SyncMap) Get(key interface{}) (interface{}, bool) {
	return m.Cache.Load(key)
}

func (m *SyncMap) Upsert(key interface{}, value interface{}, cb UpdateCb) {
	m.mu.Lock()
	defer m.mu.Unlock()

	in, ok := m.Cache.LoadOrStore(key, value)
	if ok {
		ret := cb(ok, in, value)
		m.Cache.Store(key, ret)
	}
}

func (m *SyncMap) Delete(key interface{}) {
	m.Cache.Delete(key)
}

type Tuple struct {
	Key interface{}
	Val interface{}
}

func (m *SyncMap) Foreach(f Proc) {
	ret := make([]Tuple, 0)
	m.Cache.Range(func(key, value interface{}) bool {
		ret = append(ret, Tuple{Key: key, Val: value})
		return true
	})

	for _, iter := range ret {
		f(iter.Key, iter.Val)
	}
}

type ConcurrentMap struct {
	Cache cmap.ConcurrentMap
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		Cache: cmap.New(),
	}
}

func (m ConcurrentMap) Set(key interface{}, value interface{}) {
	m.Cache.Set(fmt.Sprintf("%v", key), value)
}

func (m ConcurrentMap) Get(key interface{}) (interface{}, bool) {
	return m.Cache.Get(fmt.Sprintf("%v", key))
}

func (m ConcurrentMap) Upsert(key interface{}, value interface{}, cb UpdateCb) {
	m.Cache.Upsert(fmt.Sprintf("%v", key), value, cmap.UpsertCb(cb))
}

func (m ConcurrentMap) Delete(key interface{}) {
	m.Cache.Remove(fmt.Sprintf("%v", key))
}

func (m ConcurrentMap) Foreach(f Proc) {
	for iter := range m.Cache.IterBuffered() {
		f(iter.Key, iter.Val)
	}
}