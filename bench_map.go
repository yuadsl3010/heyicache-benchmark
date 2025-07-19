package main

import "sync"

type TestMap struct {
	c    map[string]*TestStruct
	lock sync.RWMutex
}

func (m *TestMap) Get(key string) (*TestStruct, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.c[key]
	return v, ok
}

func (m *TestMap) Set(key string, value *TestStruct) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.c[key] = value
	return nil
}
