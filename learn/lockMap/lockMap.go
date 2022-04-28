package lockMap

import (
	"log"
	"sync"
)

type MapCache struct {
	lock sync.RWMutex
	data map[string]string
}

var (
	Mcache *MapCache
)

func init() {
	log.Println("init")
	Mcache = newMapCache()
}
func newMapCache() *MapCache {
	return &MapCache{
		data: make(map[string]string),
	}
}

func (m *MapCache) Set(key, value string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
	return true
}

func (m *MapCache) Get(key string) (value string, exist bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	data, exist := m.data[key]
	if exist {
		return data, true
	}
	return "", false
}
